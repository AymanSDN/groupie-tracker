package source

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

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
