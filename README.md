# postgres-bruteforcer

This Go script reads a list of usernames and passwords from a file called `creds.txt` and tests them against a PostgreSQL database. The script accepts an IP of the PostgreSQL server as a command line argument. If an access is granted, the script writes this output to a file called `pg-output.txt`.


