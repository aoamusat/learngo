package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"syreclabs.com/go/faker"
)

// User represents a user object.
type User struct {
	ID        int
	UserAgent string
	IPAddress string
}

type Movie struct {
	Title  string
	Year   int
	Color  bool
	Actors []string
}

func main() {
	log.SetPrefix("Server Log: ")
	port := getServerPort()
	serverConfig := getConfig()
	log.Println(serverConfig)
	fmt.Println("Server running on port: ", port)
	http.HandleFunc("/", handler)
	http.HandleFunc("/movies", GetMoviesHandler)
	err := http.ListenAndServe("localhost:"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// handler echoes the Path component of the request URL r.
func handler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("X-API-RateLimit-Limit", "3600")
	responseData := make(map[string]any)
	responseData["path"] = request.URL.Path
	user := User{
		IPAddress: request.RemoteAddr,
		ID:        rand.Intn(10),
		UserAgent: request.Header.Get("User-Agent"),
	}
	responseData["user"] = user
	jsonData, err := json.Marshal(responseData)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Write(jsonData)
}

func GetMoviesHandler(response http.ResponseWriter, request *http.Request) {
	// Declare slice of movies
	var movies []Movie

	for i := 0; i < 10; i++ {
		movies = append(movies, Movie{
			Title:  faker.Lorem().Sentence(2),
			Year:   faker.RandomInt(2010, 2024),
			Actors: []string{faker.Name().Name(), faker.Name().Name(), faker.Name().Name()},
			Color:  []bool{true, false}[faker.RandomInt(0, 1)],
		})
	}

	moviesJson, err := json.Marshal(movies)

	if err != nil {
		log.Fatal(err)
	}

	response.Header().Set("Content-Type", "application/json")

	response.Write((moviesJson))
}
