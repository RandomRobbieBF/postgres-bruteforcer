package main

import (
    "bufio"
    "os"
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "log"
    "strings"
)




func connector(url string) {
	
	
	// open the creds file
	file, err := os.Open("creds.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read the creds file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		creds := strings.Split(scanner.Text(), ":")
		username := creds[0]
		password := creds[1]

		// Connect to the "postgres" database using the "postgres" username and an empty password
    db, err := sql.Open("postgres", "postgres://"+username+":"+password+"@"+url+":5432/template1?sslmode=disable")
    if err != nil {
        // If the initial connection failed, retry with "sslmode=require"
        fmt.Println("Unable to connect to the database "+url+" with SSL:", err)
        db, err = sql.Open("postgres", "postgres://"+username+":"+password+"@"+url+":5432/template1?sslmode=required")
        if err != nil {
            fmt.Println("Unable to connect to the database "+url+" with No SSL:", err)
            return
        }
    }
    defer db.Close()

    // Make sure we can actually connect to the database
    err = db.Ping()
    if err != nil {
        fmt.Println("Unable to connect to the database "+url+":", err)
        return
    }

	
	
	

    
    // Query the "pg_database" system catalog to get the names of all the databases
    rows, err := db.Query("SELECT datname FROM pg_database WHERE datistemplate = false")
    if err != nil {
        fmt.Println("error getting database details")
    }
    defer rows.Close()

    // Iterate over the rows and print the names of the databases
    for rows.Next() {
        var datname string
        if err := rows.Scan(&datname); err != nil {
            fmt.Println("error getting database details")
        }
        fmt.Println(datname)
        // Open the file in append mode
			f, err := os.OpenFile("pg-output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()

			// Write the string to the file with a new line after it
			_, err = fmt.Fprintln(f, ""+datname+" - postgres://"+username+":"+password+"@"+url+":5432")
			if err != nil {
				fmt.Println(err)
				return
			}
    }

    // Make sure we didn't miss any errors while iterating over the rows
    if err := rows.Err(); err != nil {
        panic(err)
    }
}
}
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a URL or file as an argument")
		return
	}

	input := os.Args[1]

	// Check if the argument is a file
	if _, err := os.Stat(input); err == nil {
		// Open the file
		f, err := os.Open(input)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		// Read the file line by line
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			// Set the URL variable as the scanned line
			url := scanner.Text()

			// Send the URL to the grabber function
			connector(url)
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
	} else {
		// Set the URL variable as the input argument
		url := input

		// Send the URL to the grabber function
		connector(url)
	}
}
