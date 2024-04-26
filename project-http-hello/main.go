package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", GreetingHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GreetingHandler(ResponseWriter http.ResponseWriter, Request *http.Request) {
	var name string = Request.URL.Query().Get("name")

	fmt.Fprint(ResponseWriter, "Hello, "+name)
}
