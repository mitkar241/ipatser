package main

import (
	"fmt"
    "database/sql"

	_ "github.com/lib/pq"
)

const (
	DB_IP       = "vcs_postgres" // for manual deployment use `localhost`
    DB_USER     = "raktim"
    DB_PASSWORD = "12345678"
    DB_NAME     = "movies"
)

/*
vcs_postgres | 2022-03-28 18:21:38.429 UTC [25] LOG:  database system was shut down at 2022-03-28 18:19:35 UTC
vcs_postgres | 2022-03-28 18:21:38.442 UTC [1] LOG:  database system is ready to accept connections
vcs_ipatser | 2022/03/28 18:21:50 Getting movies...
vcs_postgres | 2022-03-28 18:21:50.360 UTC [32] FATAL:  database "movies" does not exist
vcs_ipatser | 2022/03/28 18:21:50 http: panic serving 172.18.0.1:57492: pq: database "movies" does not exist
*/

// DB set up
func setupDB() *sql.DB {
    dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", DB_IP, DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    return db
}
