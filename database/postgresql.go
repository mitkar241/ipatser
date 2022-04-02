package database

import (
	"fmt"
    "database/sql"
    "strings"
    "github.com/ipatser/utils"

	_ "github.com/lib/pq"
)

const (
	DB_IP       = "localhost" // for manual deployment use `localhost`
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
    utils.CheckErr(err)
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
    utils.CheckErr(err)
    return db
}

/*
##########
# Create Database
##########
*/
func CreateDatabase(dbName string) {
    db := setupDB0()
    utils.PrintMessage("Creating Database " + dbName)
    query := "CREATE DATABASE " + dbName
    utils.PrintMessage(query)
    _, err := db.Exec(query)
    if err == nil {
        utils.PrintMessage("successfully created DB")
    } else if strings.Contains(err.Error(), "already exists") {
        utils.PrintMessage(err.Error())
    } else {
        // check errors
        utils.CheckErr(err)
    }
}

/*
##########
# Create Table
##########
*/
func CreateTableMovies() {
    db := setupDB()
    utils.PrintMessage("Creating Table movies")
    _, err := db.Exec("CREATE TABLE movies(id SERIAL, movieID varchar(50) NOT NULL, movieName varchar(50) NOT NULL, PRIMARY KEY (id))")
    if err == nil {
        utils.PrintMessage("successfully created table movies")
    } else if strings.Contains(err.Error(), "already exists") {
        utils.PrintMessage(err.Error())
    } else {
        // check errors
        utils.CheckErr(err)
    }
}
