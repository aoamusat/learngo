package main

import (
	"fmt"
	"log"
	"net/http"
)

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
	fmt.Fprintf(response, "URL.Path = %q\n", request.URL.Path)
}
