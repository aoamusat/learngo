package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, world!")
	fmt.Printf("#QOTD: %s\n", quote.Opt())
	fmt.Printf("Session ID: %s\n", generateSessionID())
}
