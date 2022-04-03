package database

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/ipatser/utils"
	_ "github.com/lib/pq"
)

/*
##########
# Global Variables used by Database
##########
*/
var (
	DB_IP       = utils.GetCfgVar("DB_IP")
	DB_USER     = utils.GetCfgVar("DB_USER")
	DB_PASSWORD = utils.GetCfgVar("DB_PASSWORD")
	DB_NAME     = "movies"
)

/*
##########
# Setup Database Initial
##########
*/
func SetupDB0() *sql.DB {
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", DB_IP, DB_USER, DB_PASSWORD, "postgres")
	db, err := sql.Open("postgres", dbinfo)
	// if connection to database `postgres` fails,
	// app must panic
	utils.PanicIfErr(err)
	return db
}

/*
##########
# Setup Database
##########
*/
func SetupDB() *sql.DB {
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", DB_IP, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	// if connection to custom database fails,
	// app must panic
	utils.PanicIfErr(err)
	return db
}

/*
##########
# Create Database
##########
*/
func CreateDatabase(dbName string) {
	db := SetupDB0()

	utils.PrintMsg("Creating Database " + dbName)
	query := "CREATE DATABASE " + dbName
	_, err := db.Exec(query)
	if err == nil {
		// successful database creation
		utils.PrintMsg("successfully created DB")
	} else if strings.Contains(err.Error(), "already exists") {
		// if database already exists,
		// app doesn't need to panic
		utils.CheckErr(err)
	} else {
		// if database is not created,
		// app must panic
		utils.PanicIfErr(err)
	}
}

/*
##########
# Create Table
##########
*/
func CreateTableMovies() {
	db := SetupDB()

	utils.PrintMsg("Creating Table movies")
	_, err := db.Exec("CREATE TABLE movies(id SERIAL, movieID varchar(50) NOT NULL, movieName varchar(50) NOT NULL, PRIMARY KEY (id))")
	if err == nil {
		// successful table creation
		utils.PrintMsg("successfully created table movies")
	} else if strings.Contains(err.Error(), "already exists") {
		// if table already exists,
		// app doesn't need to panic
		utils.CheckErr(err)
	} else {
		// if table is not created,
		// app must panic
		utils.PanicIfErr(err)
	}
}
