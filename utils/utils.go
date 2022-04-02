package utils

import (
	"log"
)

// Function for handling messages
func PrintMessage(message string) {
    log.Println(message)
}

// Function for handling errors
func CheckErr(err error) {
    if err != nil {
        panic(err)
    }
}
