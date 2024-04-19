package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func HelloGreeting(name string) (string, error) {
	if name == "" {
		return "", errors.New("name cannot be empty")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
