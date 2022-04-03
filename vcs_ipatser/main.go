package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ipatser/database"
	"github.com/ipatser/movies"
	"github.com/ipatser/utils"
)

func main() {
	port := "8000"

	database.CreateDatabase("movies")
	database.CreateTableMovies()

	// Init the mux router
	router := mux.NewRouter()

	// Route handles & endpoints

	// Get all movies
	router.HandleFunc("/movies", movies.GetMovieAll).Methods("GET")

	// Create a movie
	router.HandleFunc("/movies", movies.CreateMovie).Methods("POST")

	// Delete a specific movie by the movieID
	router.HandleFunc("/movies/{movieid}", movies.DeleteMovieById).Methods("DELETE")

	// Delete all movies
	router.HandleFunc("/movies", movies.DeleteMovieAll).Methods("DELETE")

	// serve the app
	utils.PrintMsg("Server started at " + port + "...")
	utils.PrintMsg(http.ListenAndServe(":"+port, router).Error())
}
