package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const URL_PATTERN string = "https://xkcd.com/%d/info.0.json"

type Comic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

func FetchComic(num int) (Comic, error) {
	var comic Comic
	response, err := http.Get(fmt.Sprintf(URL_PATTERN, num))

	if err != nil {
		return comic, errors.New("Failed to retrieve comic: " + err.Error())
	}

	if err := json.NewDecoder(response.Body).Decode(&comic); err != nil {
		return comic, err
	}

	defer response.Body.Close()

	return comic, nil
}

func SaveComicsToLocalIndex(index map[int]Comic, filename string) error {
	JsonData, err := json.MarshalIndent(index, "", "   ")

	if err != nil {
		return err
	}

	FileWriteErr := os.WriteFile(filename, JsonData, 0644)

	if FileWriteErr != nil {
		return FileWriteErr
	}

	return nil
}

func LoadIndexFromFile(filename string) (map[int]Comic, error) {
	ComicMap := make(map[int]Comic)

	// Read the JSON data from the file
	indexJSON, err := os.ReadFile(filename)
	if err != nil {
		return ComicMap, err
	}

	// Unmarshal the JSON data into a map
	err = json.Unmarshal(indexJSON, &ComicMap)
	if err != nil {
		return ComicMap, err
	}

	return ComicMap, nil
}

func XkcdInit() {
	fmt.Println("Initializing comic index from 1970 till date...")
	spinner(100 * time.Millisecond)
	ComicMap := make(map[int]Comic)
	for idx := 1000; idx <= 1999; idx++ {
		comic, err := FetchComic(idx)
		if err != nil {
			log.Printf("Failed to fetch comic %d: %s", idx, err.Error())
			continue
		}
		ComicMap[comic.Num] = comic
	}
	SaveComicsToLocalIndex(ComicMap, "xkcd_index.json")
	fmt.Printf("Comic index successfully initialized. :)")
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func XkcdSearch(SearchTerm string) {
	query := strings.ToLower(SearchTerm)
	XkcdIndex, err := LoadIndexFromFile("xkcd_index.json")

	if err != nil {
		fmt.Printf("Error loading local comics database: %v\nDid you run: xkcd init?", err)
		return
	}

	for _, comic := range XkcdIndex {
		if strings.Contains(strings.ToLower(comic.Transcript), query) ||
			strings.Contains(strings.ToLower(comic.Title), query) {
			fmt.Println("Comic Title:", comic.Title)
			fmt.Println("Comic Transcript:", comic.Transcript)
			fmt.Println("Comic URL:", fmt.Sprintf("https://xkcd.com/%d/", comic.Num))
			fmt.Println("-----------------------------------")
		}
	}
}
