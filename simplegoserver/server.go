package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

// User represents a user object.
type User struct {
	ID        int    `json:"id"`
	UserAgent string `json:"userAgent"`
	IPAddress string `json:"ipAddress"`
}

func main() {
	log.SetPrefix("Server Error: ")
	port := "8080"
	fmt.Println("Server running on port: ", port)
	http.HandleFunc("/", handler)
	err := http.ListenAndServe("localhost:"+port, nil)
	log.Fatal(err)
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
