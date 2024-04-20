package main

import "fmt"

func main() {
	var users [5]string

	users[4] = "John Smith"

	for i, user := range users {
		if user != "" {
			fmt.Printf("%d => %s\n\a", i+1, user)
		} else {
			continue
		}
	}
}
