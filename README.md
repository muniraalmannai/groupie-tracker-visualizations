
# 🎵 Groupie Tracker Visualisations

Groupie Tracker is a web-based application designed to help users explore and track details about musical artists. The application features a sleek, user-friendly interface for discovering artist information, including member details, creation dates, first albums, and upcoming concert dates. Users can navigate through artist profiles and enjoy a responsive design that adapts seamlessly to any device.

## Project Overview

The application allows users to:
1. View a grid of artists with key information.
2. Access detailed pages for each artist, including members, creation dates, first albums, and concert schedules.
3. Handle errors gracefully with custom error pages and navigation options.

### Features

- **Artist Directory:** An interactive grid showcasing various artists with their images and essential details.
- **Detailed Artist Profiles:** Comprehensive artist pages providing information on band members, album history, and tour locations.
- **Responsive Design:** Optimized for various screen sizes, ensuring a smooth experience on mobile and desktop devices.
- **Dynamic Data Rendering:** Utilizes Go templates for rendering artist data dynamically.
- **Error Handling:** Custom error pages for a better user experience.

## File Structure

```
├── client.go            # Handles client-side requests and API interactions
├── handlers.go          # Contains the main HTTP handlers for routing
├── types.go             # Defines data structures used across the application
├── styles.css           # Main CSS file for styling and responsive design
├── index.html           # Main page listing all artists
├── artist_details.html  # Detailed view template for individual artists
├── error.html           # Error page template
├── go.mod               # Module dependencies and package management
├── static/              # Contains static assets like images
│   └── images/
│       └── graph.png    # Image file for request flow graph
```

## Technologies Used

- **Backend:** Go (Golang) - Only Go is allowed for this project
- **Frontend:** HTML, CSS (Flexbox, Grid Layout)
- **Styling:** Custom CSS with theming and responsive design
- **Templating:** Go's `html/template` for server-side rendering

### Request Flow Diagram

Below is a visual representation of the route flow for a user request in the Groupie Tracker application:

![Request Flow Diagram](static/images/graph.png)

### Explanation

1. **User Request**: The user makes a request (e.g., clicks an artist card).
2. **Go Backend**: The request is received and processed by the backend.
3. **`client.go (API Fetch)`**: The backend fetches data from an external API.
4. **API**: The data is retrieved and sent back to `client.go`.
5. **`handlers.go (Data Processing)`**: Processes the data, utilizing data structures defined in **`types.go`**.
6. **Template Rendering**: The processed data is rendered into an HTML template (`index.html`, `artist_details.html`, or `error.html`).
7. **Browser (Display)**: The rendered HTML is sent back to the user’s browser for display.

### Installation

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd groupie-tracker
   ```

2. Install Go dependencies:

   ```bash
   go mod tidy
   ```

3. Run the application locally:

   ```bash
   go run main.go
   ```

4. Open the browser and go to [http://localhost:8080](http://localhost:8080) to access the application.

### Usage

- Navigate the main page to see a list of artists.
- Click on an artist card for more detailed information.
- Use the back button or navigation links to return to the main page.

### Example Input

The application retrieves data dynamically, but the structure of the artist data includes:
- **Name:** Artist or band name
- **Creation Date:** Year the artist or band was formed
- **First Album:** Name and release date of the first album
- **Members:** List of band members
- **Concert Dates:** Locations and dates of upcoming concerts

### Example Output

The application displays a list of artists, and on selecting an artist, the detailed profile includes:
- **Image:** Artist or band photo
- **Information:** Creation date, first album, and members
- **Concert Schedule:** List of upcoming tour locations and dates

## Development Notes

This project was developed with a focus on clean code and modular design. The Go backend handles routing and data rendering, while the frontend utilizes modern CSS techniques for a polished user experience. Custom Go types and functions ensure a robust and scalable codebase.

### Known Issues

- **Search Functionality:** Future improvements may include implementing a search feature for easier navigation.
- **Data Source:** Currently, the data is static but could be integrated with an API for real-time updates.

## License

This project was completed as part of a coursework assignment at Reboot01.
