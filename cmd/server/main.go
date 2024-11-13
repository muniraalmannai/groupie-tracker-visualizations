package main

import (
	"groupie-trackers/pkg/api"
	"groupie-trackers/pkg/handlers"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type TemplateData struct {
	Artists      []api.Artist
	Query        string
	IsSearchPage bool // Add this field
}

func main() {
	mux := http.NewServeMux()

	// Handle the main page with preloaded artists
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet { // Only allow GET requests
			handlers.RenderErrorPage(w, "Method Not Allowed", "The method is not allowed for this resource.", http.StatusMethodNotAllowed)
			return
		}

		if r.URL.Path != "/" {
			handlers.RenderErrorPage(w, "Page Not Found", "The page you are looking for does not exist.", http.StatusNotFound)
			return
		}

		// Fetch artists data to display on the main page
		artists, err := api.FetchArtists("https://groupietrackers.herokuapp.com/api/artists")
		// Search bar functionality
		// Filter artists if a search query is present
		query := r.URL.Query().Get("search")
		if query != "" {
			filteredArtists := []api.Artist{}
			for _, artist := range artists {
				// Search in artist name
				if containsIgnoreCase(artist.Name, query) {
					filteredArtists = append(filteredArtists, artist)
					continue
				}
				// Search in members
				for _, member := range artist.Members {
					if containsIgnoreCase(member, query) {
						filteredArtists = append(filteredArtists, artist)
						break
					}
				}
				// Search in locations
				// for _, location := range artist.LocationDetails {
				// 	if containsIgnoreCase(location, query) {
				// 		filteredArtists = append(filteredArtists, artist)
				// 		break
				// 	}
				// }
			}
			artists = filteredArtists
		}
		if err != nil {
			handlers.RenderErrorPage(w, "Error", "Failed to load artists. Please try again later.", http.StatusInternalServerError)
			return
		}

		// Render the main page with the artists data
		tmpl, err := template.ParseFiles("web/templates/index.html")
		if err != nil { // Handle missing or misnamed template file
			log.Printf("Error parsing index.html: %v", err)
			handlers.RenderErrorPage(w, "Internal Server Error", "An error occurred while loading the page. Please try again later.", http.StatusInternalServerError)
			return
		}
		data := TemplateData{
			Artists:      artists,
			Query:        query,
			IsSearchPage: query != "",
		}

		err = tmpl.Execute(w, data) // Pass data to the template
		if err != nil {
			handlers.RenderErrorPage(w, "Internal Server Error", "An error occurred while rendering the page. Please try again later.", http.StatusInternalServerError)
			log.Printf("Template execution error: %v", err)
		}
	})

	// Handle artist details pages
	mux.HandleFunc("/artists/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet { // Only allow GET requests
			handlers.RenderErrorPage(w, "Method Not Allowed", "The method is not allowed for this resource.", http.StatusMethodNotAllowed)
			return
		}

		idStr := r.URL.Path[len("/artists/"):]
		id, err := strconv.Atoi(idStr)
		if err != nil || idStr == "" {
			handlers.RenderErrorPage(w, "Invalid Artist ID", "The artist ID provided is not valid.", http.StatusBadRequest)
			return
		}
		if handlers.HandleArtistByID(w, r, id) {
			// Error already handled in HandleArtistByID, no need to handle it again
			return
		}
	})

	// Handle invalid paths for /artists
	mux.HandleFunc("/artists", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet { // Only allow GET requests
			handlers.RenderErrorPage(w, "Method Not Allowed", "The method is not allowed for this resource.", http.StatusMethodNotAllowed)
			return
		}

		handlers.RenderErrorPage(w, "Page Not Found", "The page you are looking for does not exist.", http.StatusNotFound)
	})

	// Serve static files (CSS, JS, etc.)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static/"))))

	// Start the server
	log.Println("Server is running on port 8080...")
	log.Println("Server is accessible at: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

// Utility function to check if a string contains another string, case-insensitively
func containsIgnoreCase(str, substr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}
