package main 

import (
    "fmt"
    "log"
    "net/http"
    "github.com/ipatser/database"
    "github.com/ipatser/movies"

    "github.com/gorilla/mux"
)

func main() {
    port := "8000"

    database.CreateDatabase("movies")
    database.CreateTableMovies()

    // Init the mux router
    router := mux.NewRouter()

    // Route handles & endpoints

    // Get all movies
    router.HandleFunc("/movies", movies.GetMovies).Methods("GET")

    // Create a movie
    router.HandleFunc("/movies", movies.CreateMovie).Methods("POST")

    // Delete a specific movie by the movieID
    router.HandleFunc("/movies/{movieid}", movies.DeleteMovie).Methods("DELETE")

    // Delete all movies
    router.HandleFunc("/movies", movies.DeleteMovies).Methods("DELETE")

    // serve the app
    fmt.Println("Server at " + port + "...")
    log.Fatal(http.ListenAndServe(":" + port, router))
}
