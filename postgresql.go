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
##########
# Setup Database Initial
##########
*/
func setupDB0() *sql.DB { 
    dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", DB_IP, DB_USER, DB_PASSWORD, "postgres")
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    return db
}

/*
##########
# Setup Database
##########
*/
func setupDB() *sql.DB { 
    dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", DB_IP, DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    return db
}

/*
##########
# Create Database
##########
*/
func CreateDatabase(dbName string) {
    db := setupDB0()
    printMessage("Creating Database " + dbName)
    query := "CREATE DATABASE " + dbName
    printMessage(query)
    _, err := db.Exec(query)
    // check errors
    checkErr(err)
    printMessage("successfully created DB")
}

/*
##########
# Create Table
##########
*/
func CreateTableMovies() {
    db := setupDB()
    printMessage("Creating Table movies")
    _, err := db.Exec("CREATE TABLE movies(id SERIAL, movieID varchar(50) NOT NULL, movieName varchar(50) NOT NULL, PRIMARY KEY (id))")
    // check errors
    checkErr(err)
}
