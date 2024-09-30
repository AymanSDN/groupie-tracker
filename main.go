package main

import (
	"fmt"
	"net/http"

	"groupie-tracker/source"
)

const Host, Port = "localhost", ":8080"

// Main function to start the server
func main() {
	source.FetchData() // Load the data from the API

	// Set up HTTP handler
	http.HandleFunc("/", source.HomePageHandler)
	http.HandleFunc("/artist", source.ArtistDetailsHandler)

	// Start the server
	fmt.Println("Server starting on port " + Port + "...")
	fmt.Println("ctrl + click to open: http://" + Host + Port + "/")
	http.ListenAndServe(Host+Port, nil)
}
