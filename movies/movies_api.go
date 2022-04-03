package movies

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ipatser/database"
	"github.com/ipatser/utils"
)

/*
##########
# Maintain Single Database Connectivity
##########
*/

var DB_CONN *sql.DB

func init() {
	DB_CONN = database.SetupDB()
}

/*
##########
# Get all movies
##########
*/
func GetMovieAll(res http.ResponseWriter, req *http.Request) {
	var response = JsonResponse{}
	var err error
	var movies []Movie

	utils.PrintMsg("Getting movies...")
	// Get all entries from movies table `movies``
	rows, err := DB_CONN.Query("SELECT * FROM movies")

	// if there is error getting entries from table movies
	if err != nil {
		utils.CheckErr(err)
		response = JsonResponse{Type: "error", Message: "Error while getting entries from table movies : " + err.Error()}
		json.NewEncoder(res).Encode(response)
		return
	}

	// For each entries in table `movies``
	for rows.Next() {
		// variables present in the entries in table `movies`
		var (
			id        int
			movieID   string
			movieName string
		)

		err = rows.Scan(&id, &movieID, &movieName)

		// if there is error scanning entries from table movies
		if err != nil {
			utils.CheckErr(err)
			response = JsonResponse{Type: "error", Message: "Error while scanning entries from table movies : " + err.Error()}
			json.NewEncoder(res).Encode(response)
			return
		}

		// on successful scanning of entry from table movies
		utils.PrintMsg("scanned entry with movieID: " + movieID + " and name: " + movieName + " successfully")

		movies = append(movies, Movie{MovieID: movieID, MovieName: movieName})
	}

	// on successful scanning of entry from table movies
	utils.PrintMsg("got all entries from table movies successfully")

	response = JsonResponse{Type: "success", Data: movies}
	json.NewEncoder(res).Encode(response)
}

/*
##########
# Create a movie
##########
*/
func CreateMovie(res http.ResponseWriter, req *http.Request) {
	var response = JsonResponse{}
	var err error
	var lastInsertID int

	movieID := req.FormValue("movieid")
	movieName := req.FormValue("moviename")

	// if `movieID` or `movieName` is empty,
	// the entry is invalid
	if movieID == "" || movieName == "" {
		response = JsonResponse{Type: "error", Message: "You are missing movieID or movieName parameter."}
		json.NewEncoder(res).Encode(response)
		return
	}

	utils.PrintMsg("Inserting new movie with ID: " + movieID + " and name: " + movieName)
	err = DB_CONN.QueryRow("INSERT INTO movies(movieID, movieName) VALUES($1, $2) returning id;", movieID, movieName).Scan(&lastInsertID)

	// if there is error inserting entry into table movies
	if err != nil {
		utils.CheckErr(err)
		response = JsonResponse{Type: "error", Message: "Error while inserting movie into table movies : " + err.Error()}
		json.NewEncoder(res).Encode(response)
		return
	}

	// on successful insertion of entry into table movies
	utils.PrintMsg("Inserted new movie with ID: " + movieID + " and name: " + movieName + " successfully")

	response = JsonResponse{Type: "success", Message: "The movie has been inserted successfully!"}
	json.NewEncoder(res).Encode(response)
}

/*
##########
# Delete a movie by ID
##########
*/
func DeleteMovieById(res http.ResponseWriter, req *http.Request) {
	var response = JsonResponse{}
	var err error

	params := mux.Vars(req)
	movieID := params["movieid"]

	// if `movieID` provided in empty
	// entry from table can't be deleted
	if movieID == "" {
		response = JsonResponse{Type: "error", Message: "You are missing movieID parameter."}
		json.NewEncoder(res).Encode(response)
		return
	}

	utils.PrintMsg("Deleting movie from table movies with movieID = " + movieID)
	_, err = DB_CONN.Exec("DELETE FROM movies where movieID = $1", movieID)

	// if there is error deleting entry by movieID from table movies
	if err != nil {
		utils.CheckErr(err)
		response = JsonResponse{Type: "error", Message: "Error while deleting movie by ID : " + err.Error()}
		json.NewEncoder(res).Encode(response)
		return
	}

	// on successful deletion of entry by movieID from table movies
	utils.PrintMsg("Deleted movie from table movies with movieID = " + movieID + " successfully")

	response = JsonResponse{Type: "success", Message: "The movie has been deleted successfully!"}
	json.NewEncoder(res).Encode(response)
}

/*
##########
# Delete all movies
##########
*/
func DeleteMovieAll(res http.ResponseWriter, req *http.Request) {
	var response = JsonResponse{}
	var err error

	utils.PrintMsg("Deleting all movies...")
	_, err = DB_CONN.Exec("DELETE FROM movies")

	// if there is error deleting all entries from table movies
	if err != nil {
		utils.CheckErr(err)
		response = JsonResponse{Type: "error", Message: "Error while deleting all entries from table movies : " + err.Error()}
		json.NewEncoder(res).Encode(response)
		return
	}

	// on successful deletion of all entries from table movies
	utils.PrintMsg("All movies have been deleted successfully!")

	response = JsonResponse{Type: "success", Message: "All movies have been deleted successfully!"}
	json.NewEncoder(res).Encode(response)
}
