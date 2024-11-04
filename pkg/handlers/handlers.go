package handlers

import (
	"encoding/json"
	"fmt"
	"groupie-trackers/pkg/api"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Function to render the custom error page with a status code
func RenderErrorPage(w http.ResponseWriter, title string, message string, statusCode int) {
	w.WriteHeader(statusCode) // Set the HTTP status code

	tmpl, err := template.ParseFiles("web/templates/error.html")
	if err != nil {
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		log.Printf("Error parsing template: %v", err)
		return
	}

	data := struct {
		Title   string
		Message string
	}{
		Title:   title,
		Message: message,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}

func HandleArtists(w http.ResponseWriter, r *http.Request) {
	// Check if the method is GET
	if r.Method != http.MethodGet {
		RenderErrorPage(w, "Method Not Allowed", "The method is not allowed for this resource.", http.StatusMethodNotAllowed)
		return
	}

	apiURL := "https://groupietrackers.herokuapp.com/api/artists"
	log.Printf("Fetching artists data from URL: %s", apiURL)

	artists, err := api.FetchArtists(apiURL)
	if err != nil {
		log.Printf("Error fetching artists: %v", err)
		RenderErrorPage(w, "Error", "Failed to fetch artists. Please try again later.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artists)
}

func HandleArtistByID(w http.ResponseWriter, r *http.Request, id int) bool { // Return true if an error occurred
	// Check if the method is GET
	if r.Method != http.MethodGet {
		RenderErrorPage(w, "Method Not Allowed", "The method is not allowed for this resource.", http.StatusMethodNotAllowed)
		return true
	}

	apiURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", id)
	log.Printf("Fetching artist data from URL: %s", apiURL)

	artist, err := api.FetchArtistByID(apiURL)
	if err != nil || artist == nil {
		log.Printf("Artist not found or error fetching artist: %v", err)
		RenderErrorPage(w, "Artist Not Found", "The artist you are looking for does not exist.", http.StatusNotFound)
		return true // Error occurred
	}

	tmpl, err := template.ParseFiles("web/templates/artist_details.html")
	if err != nil { // Handle missing or misnamed template file
		log.Printf("Error parsing artist_details.html: %v", err)
		RenderErrorPage(w, "Internal Server Error", "An error occurred while loading the page. Please try again later.", http.StatusInternalServerError)
		return true
	}

	err = tmpl.Execute(w, artist)
	if err != nil {
		RenderErrorPage(w, "Internal Server Error", "An error occurred while rendering the artist details. Please try again later.", http.StatusInternalServerError)
		log.Printf("Template execution error: %v", err)
		return true // Error occurred
	}

	return false // No error
}

// New handler for invalid paths
func Handle404(w http.ResponseWriter, r *http.Request) {
	// Check if the method is GET
	if r.Method != http.MethodGet {
		RenderErrorPage(w, "Method Not Allowed", "The method is not allowed for this resource.", http.StatusMethodNotAllowed)
		return
	}

	RenderErrorPage(w, "Page Not Found", "The page you are looking for does not exist.", http.StatusNotFound)
}

// Middleware to check file existence
func CheckFileExists(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the method is GET
		if r.Method != http.MethodGet {
			RenderErrorPage(w, "Method Not Allowed", "The method is not allowed for this resource.", http.StatusMethodNotAllowed)
			return
		}

		path := "." + r.URL.Path
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			RenderErrorPage(w, "File Not Found", "The file you are looking for is missing or has been renamed.", http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}
