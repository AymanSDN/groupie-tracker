	package main

	import (
		"fmt"
		"html/template"
		"log"
		"net/http"
		"path/filepath"
		"strconv"

		"groupie-tracker/source"
	)

	// Function to render the template
	func HomePageHandler(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, "404 Page Not Found", http.StatusNotFound)
			return
		}

		if r.Method != "GET" {
			http.Error(w, "methode not allowed", http.StatusMethodNotAllowed)
			return
		}

		tmpl, err := template.ParseFiles(filepath.Join("./templates/index.html"))
		if err != nil {
			log.Fatalf("Error parsing template: %v", err)
		}

		if err := tmpl.Execute(w, source.ArtistProfiles); err != nil {
			log.Printf("Error executing template: %v", err)
			http.Error(w, "error executing template", http.StatusInternalServerError)
		}
	}

	func ArtistDetailsHandler(w http.ResponseWriter, r *http.Request) {
		// if
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "Bad Request: invalid ID", http.StatusBadRequest)
			return
		}
		if id > 52 || id < 1 {
			http.Error(w, "Artist not found", http.StatusNotFound)
			return
		}
		artistData := source.LoadArtistInfos(id)
		// Parse and execute the template
		tmpl, err := template.ParseFiles(filepath.Join("./templates/artist-details.html"))
		if err != nil {
			log.Fatalf("Error parsing template: %v", err)
		}

		// Pass the artist data to the template
		if err := tmpl.Execute(w, artistData); err != nil {
			log.Printf("Error executing template: %v", err)
			http.Error(w, "Error executing template", http.StatusInternalServerError)
		}
	}
	const Host, Port = "localhost", ":8080"

	// Main function to start the server
	func main() {
		source.FetchData() // Load the data from the URL

		// Set up HTTP handler
		http.HandleFunc("/", HomePageHandler)
		http.HandleFunc("/artist", ArtistDetailsHandler)

		// Start the server
		fmt.Println("Server starting on port "+Port+"...")
		fmt.Println("ctrl + click to open: http://"+Host+Port+"/")
		http.ListenAndServe(Host+Port, nil)
	}
