package source

import (
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"text/template"
)

// Structs for the JSON data
type Artist struct {
	ID           int      `json:"id"`
	ImageUrl     string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type ArtistLocations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type ArtistDates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}
type ArtistRelation struct {
	ID    int      `json:"id"`
	Relation map[string][]string `json:"datesLocations"`
}

type ArtistData struct {
	ArtistInfos []Artist
	Dates       []ArtistDates
	Locations   []ArtistLocations
	Relations   []ArtistRelation
}

var ArtistProfiles ArtistData

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

	tmpl, err := template.ParseFiles(filepath.Join("./static/index.html"))
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	if err := tmpl.Execute(w, ArtistProfiles); err != nil {
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
	artistData := LoadArtistInfos(id)
	// Parse and execute the template
	tmpl, err := template.ParseFiles(filepath.Join("./static/artist-details.html"))
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	// Pass the artist data to the template
	if err := tmpl.Execute(w, artistData); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
