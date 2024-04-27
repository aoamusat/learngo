package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Movie struct {
	Title  string `json:"title"`
	Year   string `json:"year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"type"`
	Poster string `json:"poster"`
}

type PosterResult struct {
	Search       []*Movie `json:"Search"`
	TotalResults string   `json:"totalResults"`
	Response     string   `json:"Response"`
}

// Generate filename for movie poster
func (movie *Movie) GetFileName() string {
	var _temp []string = strings.Split(movie.Poster, ".")
	FileExt := _temp[len(_temp)-1]
	var filename string = fmt.Sprintf("%s-poster.%s", movie.ImdbID, FileExt)
	return filename
}

func (movie *Movie) SavePoster() error {
	response, err := http.Get(movie.Poster)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Create a new file to save the image
	file, err := os.Create(movie.GetFileName())
	if err != nil {
		return err
	}
	defer file.Close()

	// Copy the image data from the HTTP response body to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: poster <movie name>")
		return
	}

	MovieName := os.Args[1]

	ApiKey := os.Getenv("OMDB_API_KEY")

	query := url.QueryEscape(MovieName)
	Url := fmt.Sprintf("https://www.omdbapi.com/?apikey=%s&s=%s", ApiKey, query)
	response, err := http.Get(Url)

	if err != nil {
		fmt.Printf("error getting: %v\n", err)
	}

	defer response.Body.Close()

	var posters PosterResult

	if err := json.NewDecoder(response.Body).Decode(&posters); err != nil {
		fmt.Printf("error getting: %v\n", err)
	}

	for _, poster := range posters.Search {
		fmt.Printf("Saving poster for: %s\n", poster.Title)
		poster.SavePoster()
	}
}
