package main

import (
	"fmt"
	"log"

	"learngo.io/greetings"
	"rsc.io/quote"
)

func main() {

	log.SetPrefix("Error: ")

	message, err := greetings.HelloGreeting("Akeem")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	fmt.Printf("\n#QOTD: %s\n", quote.Go())
	fmt.Printf("Session ID: %s\n", generateSessionID())
}
