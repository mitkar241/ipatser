package main

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

// ##########
// Get all movies
// ##########

// response and request handlers
func GetMovies(w http.ResponseWriter, r *http.Request) {
    db := setupDB()
    printMessage("Getting movies...")
    // Get all movies from movies table that don't have movieID = "1"
    rows, err := db.Query("SELECT * FROM movies")
    // check errors
    checkErr(err)
    // var response []JsonResponse
    var movies []Movie
    // Foreach movie
    for rows.Next() {
        var id int
        var movieID string
        var movieName string
        err = rows.Scan(&id, &movieID, &movieName)
        // check errors
        checkErr(err)
        movies = append(movies, Movie{MovieID: movieID, MovieName: movieName})
    }
    var response = JsonResponse{Type: "success", Data: movies}
    json.NewEncoder(w).Encode(response)
}

// ##########
// Create a movie
// ##########

// response and request handlers
func CreateMovie(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("%+v\n", r)
    movieID := r.FormValue("movieid")
    movieName := r.FormValue("moviename")
    var response = JsonResponse{}
    if movieID == "" || movieName == "" {
        response = JsonResponse{Type: "error", Message: "You are missing movieID or movieName parameter."}
    } else {
        db := setupDB()
        printMessage("Inserting movie into DB")
        fmt.Println("Inserting new movie with ID: " + movieID + " and name: " + movieName)
        var lastInsertID int
    err := db.QueryRow("INSERT INTO movies(movieID, movieName) VALUES($1, $2) returning id;", movieID, movieName).Scan(&lastInsertID)
    // check errors
    checkErr(err)
    response = JsonResponse{Type: "success", Message: "The movie has been inserted successfully!"}
    }
    json.NewEncoder(w).Encode(response)
}

// ##########
// Delete a movie
// ##########

// response and request handlers
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    movieID := params["movieid"]
    var response = JsonResponse{}

    if movieID == "" {
        response = JsonResponse{Type: "error", Message: "You are missing movieID parameter."}
    } else {
        db := setupDB()
        printMessage("Deleting movie from DB")
        _, err := db.Exec("DELETE FROM movies where movieID = $1", movieID)
        // check errors
        checkErr(err)
        response = JsonResponse{Type: "success", Message: "The movie has been deleted successfully!"}
    }
    json.NewEncoder(w).Encode(response)
}

// ##########
// Delete all movies
// ##########

// response and request handlers
func DeleteMovies(w http.ResponseWriter, r *http.Request) {
    db := setupDB()
    printMessage("Deleting all movies...")
    _, err := db.Exec("DELETE FROM movies")
    // check errors
    checkErr(err)
    printMessage("All movies have been deleted successfully!")
    response := JsonResponse{Type: "success", Message: "All movies have been deleted successfully!"}
    json.NewEncoder(w).Encode(response)
}
