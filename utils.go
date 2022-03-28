package main

import (
	"log"
)

// Function for handling messages
func printMessage(message string) {
    log.Println(message)
}

// Function for handling errors
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
