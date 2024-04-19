package main

import (
	"fmt"
	"log"

	"learngo.io/greetings"
	"rsc.io/quote"
)

func main() {

	log.SetPrefix("Error: ")

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	for _, m := range messages {
		fmt.Println(m)
	}

	fmt.Printf("\n#QOTD: %s\n", quote.Go())
	fmt.Printf("Session ID: %s\n", generateSessionID())
}
