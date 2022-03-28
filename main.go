package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

// Main function
func main() {
    port := "8000"

    // Init the mux router
    router := mux.NewRouter()

    // Route handles & endpoints

    // Get all movies
    router.HandleFunc("/movies", GetMovies).Methods("GET")

    // Create a movie
    router.HandleFunc("/movies", CreateMovie).Methods("POST")

    // Delete a specific movie by the movieID
    router.HandleFunc("/movies/{movieid}", DeleteMovie).Methods("DELETE")

    // Delete all movies
    router.HandleFunc("/movies", DeleteMovies).Methods("DELETE")

    // serve the app
    fmt.Println("Server at " + port + "...")
    log.Fatal(http.ListenAndServe(":" + port, router))
}
