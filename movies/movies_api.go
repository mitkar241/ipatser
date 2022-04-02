package movies

import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/ipatser/database"
    "github.com/ipatser/utils"

    "github.com/gorilla/mux"
)

/*
##########
# Get all movies
##########
*/
func GetMovies(res http.ResponseWriter, req *http.Request) {
    db := database.SetupDB()
    utils.PrintMessage("Getting movies...")
    // Get all movies from movies table that don't have movieID = "1"
    rows, err := db.Query("SELECT * FROM movies")
    // check errors
    utils.CheckErr(err)
    // var response []JsonResponse
    var movies []Movie
    // Foreach movie
    for rows.Next() {
        var id int
        var movieID string
        var movieName string
        err = rows.Scan(&id, &movieID, &movieName)
        // check errors
        utils.CheckErr(err)
        movies = append(movies, Movie{MovieID: movieID, MovieName: movieName})
    }
    var response = JsonResponse{Type: "success", Data: movies}
    json.NewEncoder(res).Encode(response)
}

/*
##########
# Create a movie
##########
*/
func CreateMovie(res http.ResponseWriter, req *http.Request) {
    fmt.Printf("%+v\n", req)
    movieID := req.FormValue("movieid")
    movieName := req.FormValue("moviename")
    var response = JsonResponse{}
    if movieID == "" || movieName == "" {
        response = JsonResponse{Type: "error", Message: "You are missing movieID or movieName parameter."}
    } else {
        db := database.SetupDB()
        utils.PrintMessage("Inserting movie into DB")
        fmt.Println("Inserting new movie with ID: " + movieID + " and name: " + movieName)
        var lastInsertID int
        err := db.QueryRow("INSERT INTO movies(movieID, movieName) VALUES($1, $2) returning id;", movieID, movieName).Scan(&lastInsertID)
        // check errors
        utils.CheckErr(err)
        response = JsonResponse{Type: "success", Message: "The movie has been inserted successfully!"}
    }
    json.NewEncoder(res).Encode(response)
}

/*
##########
# Delete a movie
##########
*/
func DeleteMovie(res http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    movieID := params["movieid"]
    var response = JsonResponse{}

    if movieID == "" {
        response = JsonResponse{Type: "error", Message: "You are missing movieID parameter."}
    } else {
        db := database.SetupDB()
        utils.PrintMessage("Deleting movie from DB")
        _, err := db.Exec("DELETE FROM movies where movieID = $1", movieID)
        // check errors
        utils.CheckErr(err)
        response = JsonResponse{Type: "success", Message: "The movie has been deleted successfully!"}
    }
    json.NewEncoder(res).Encode(response)
}

/*
##########
# Delete all movies
##########
*/
func DeleteMovies(res http.ResponseWriter, req *http.Request) {
    db := database.SetupDB()
    utils.PrintMessage("Deleting all movies...")
    _, err := db.Exec("DELETE FROM movies")
    // check errors
    utils.CheckErr(err)
    utils.PrintMessage("All movies have been deleted successfully!")
    response := JsonResponse{Type: "success", Message: "All movies have been deleted successfully!"}
    json.NewEncoder(res).Encode(response)
}
