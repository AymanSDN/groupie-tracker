package source

import "strings"

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

type ArtistData struct {
	ArtistInfos []Artist
	Dates       []ArtistDates
	Locations   []ArtistLocations
}

var ArtistProfiles ArtistData

// Function to fetch data from the URLs
func FetchData() {
	// Load artist data from the API
	DecodeJSONFromURL("https://groupietrackers.herokuapp.com/api/artists", &ArtistProfiles.ArtistInfos)

	// Load artist dates from the API
	DecodeJSONFromURL("https://groupietrackers.herokuapp.com/api/dates", &ArtistProfiles.Dates)

	// Load artist locations from the API
	DecodeJSONFromURL("https://groupietrackers.herokuapp.com/api/locations", &ArtistProfiles.Locations)
}

func LoadArtistInfos(id int) interface{} {
	type ArtistDetails struct {
		Artist    Artist
		Locations []string
		Dates     []string
	}
	// Find the artist with the matching ID
	var selectedArtist Artist
	for _, artist := range ArtistProfiles.ArtistInfos {
		if artist.ID == id {
			artist.FirstAlbum = strings.ReplaceAll(artist.FirstAlbum, "-", "	/	")
			selectedArtist = artist
			break
		}
	}

	// Fetch corresponding locations and dates
	var artistLocations, artistDates []string

	for _, loc := range ArtistProfiles.Locations {
		if loc.ID == id {
			for i := 0; i < len(loc.Locations); i++ {
				loc.Locations[i] = strings.ReplaceAll(loc.Locations[i], "-", " ")
				artistLocations = append(artistLocations, strings.ReplaceAll(loc.Locations[i], "_", " "))
			}
			break
		}
	}

	for _, date := range ArtistProfiles.Dates {
		if date.ID == id {
			for i := 0; i < len(date.Dates); i++ {
				date.Dates[i] = strings.ReplaceAll(date.Dates[i], "-", "	/	")
				artistDates = append(artistDates, strings.ReplaceAll(date.Dates[i], "*", ""))
			}
			// artistDates = date.Dates
			break
		}
	}

	artistData := ArtistDetails{
		Artist:    selectedArtist,
		Locations: artistLocations,
		Dates:     artistDates,
	}
	return artistData
}
