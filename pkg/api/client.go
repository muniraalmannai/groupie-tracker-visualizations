package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func FetchArtists(apiURL string) ([]Artist, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from API: %v", err)
	}
	defer resp.Body.Close()

	log.Printf("HTTP GET request to %s returned status code: %d", apiURL, resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var artists []Artist
	if err := json.Unmarshal(body, &artists); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return artists, nil
}
func FetchArtistByID(apiURL string) (*Artist, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from API: %v", err)
	}
	defer resp.Body.Close()

	log.Printf("HTTP GET request to %s returned status code: %d", apiURL, resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var artist Artist
	if err := json.Unmarshal(body, &artist); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	// Fetch location details
	locationData, err := fetchLocationData(artist.Locations)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch location data: %v", err)
	}

	// Fetch concert date details
	concertDateData, err := fetchConcertDateData(artist.ConcertDates)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch concert date data: %v", err)
	}

	// Store the original order of locations
	artist.LocationDetails = locationData.Locations
	artist.LocationConcerts = associateLocationsWithDates(locationData.Locations, concertDateData.Dates)

	return &artist, nil
}

// Helper function to associate locations with their respective dates
func associateLocationsWithDates(locations []string, dates []string) map[string][]string {
	locationConcertMap := make(map[string][]string)
	currentLocationIndex := -1

	for _, date := range dates {
		if date[0] == '*' {
			// Date with asterisk, move to the next location
			currentLocationIndex++
			if currentLocationIndex < len(locations) {
				date = date[1:] // Remove the asterisk
			}
		}
		if currentLocationIndex < len(locations) {
			location := locations[currentLocationIndex]
			locationConcertMap[location] = append(locationConcertMap[location], date)
		}
	}

	return locationConcertMap
}

func fetchLocationData(url string) (*LocationData, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from API: %v", err)
	}
	defer resp.Body.Close()

	var locationData LocationData
	if err := json.NewDecoder(resp.Body).Decode(&locationData); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return &locationData, nil
}

func fetchConcertDateData(url string) (*ConcertDateData, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from API: %v", err)
	}
	defer resp.Body.Close()

	var concertDateData ConcertDateData
	if err := json.NewDecoder(resp.Body).Decode(&concertDateData); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return &concertDateData, nil
}
