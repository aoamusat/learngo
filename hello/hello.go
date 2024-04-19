package main

import (
	"fmt"

	"learngo.io/greetings"
	"rsc.io/quote"
)

func main() {

	var message = greetings.HelloGreeting("Akeem")
	fmt.Println(message)

	fmt.Printf("\n#QOTD: %s\n", quote.Go())
	fmt.Printf("Session ID: %s\n", generateSessionID())
}
