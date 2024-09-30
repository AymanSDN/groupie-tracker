package source

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Function to fetch data from the URLs
func FetchData() {
	// Load artist data from the API
	DecodeJSONFromURL("https://groupietrackers.herokuapp.com/api/artists", &ArtistProfiles.ArtistInfos)

	// Load artist dates from the API
	DecodeJSONFromURL("https://groupietrackers.herokuapp.com/api/dates", &ArtistProfiles.Dates)

	// Load artist locations from the API
	DecodeJSONFromURL("https://groupietrackers.herokuapp.com/api/locations", &ArtistProfiles.Locations)
	
	// Load artist Relations from the API
	DecodeJSONFromURL("https://groupietrackers.herokuapp.com/api/relation", &ArtistProfiles.Relations)
}

// Helper function to decode JSON from a URL
func DecodeJSONFromURL(url string, data interface{}) {
	// Make a GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching URL %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	// Check for a successful response
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: Received non-200 response code from %s: %s\n", url, resp.Status)
		return
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body from %s: %v\n", url, err)
		return
	}
	if url[42:] != "artists" {
		body = body[9 : len(body)-2]
	}
	// Decode the JSON from the response body
	decoder := json.NewDecoder(strings.NewReader(string(body)))
	if err := decoder.Decode(data); err != nil {
		fmt.Printf("Error decoding JSON from URL %s: %v\n", url, err)
	}
}
func LoadArtistInfos(id int) interface{} {
	type ArtistDetails struct {
		Artist    Artist
		Locations []string
		Dates     []string
		Relations map[string][]string
	}
	var selectedArtist Artist
	for _, artist := range ArtistProfiles.ArtistInfos {
		if artist.ID == id {
			artist.FirstAlbum = strings.ReplaceAll(artist.FirstAlbum, "-", " / ")
			selectedArtist = artist
			break
		}
	}

	var artistLocations, artistDates []string
	artistRelations := make(map[string][]string)

	// Find locations for the artist
	for _, loc := range ArtistProfiles.Locations {
		if loc.ID == id {
			for i := 0; i < len(loc.Locations); i++ {
				loc.Locations[i] = strings.ReplaceAll(loc.Locations[i], "-", " ")
				artistLocations = append(artistLocations, strings.ReplaceAll(loc.Locations[i], "_", " "))
			}
			break
		}
	}

	// Find dates for the artist
	for _, date := range ArtistProfiles.Dates {
		if date.ID == id {
			for i := 0; i < len(date.Dates); i++ {
				date.Dates[i] = strings.ReplaceAll(date.Dates[i], "-", " / ")
				artistDates = append(artistDates, strings.ReplaceAll(date.Dates[i], "*", ""))
			}
			break
		}
	}

	// Find relations for the artist
	for _, relation := range ArtistProfiles.Relations {
		if relation.ID == id {
			for location, dates := range relation.Relation {
				artistRelations[location] = dates
			}
			break
		}
	}

	// Construct the full artist details
	artistData := ArtistDetails{
		Artist:    selectedArtist,
		Locations: artistLocations,
		Dates:     artistDates,
		Relations: artistRelations,
	}
	return artistData
}
