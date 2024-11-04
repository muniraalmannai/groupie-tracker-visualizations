package api

type LocationData struct {
	Locations []string `json:"locations"`
}

type ConcertDateData struct {
	Dates []string `json:"dates"`
}

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`    // URL as a string
	ConcertDates string   `json:"concertDates"` // URL as a string
	Relations    string   `json:"relations"`

	LocationDetails  []string            `json:"-"` // List of locations in order
	LocationConcerts map[string][]string `json:"-"` // Map of location to its dates
}
