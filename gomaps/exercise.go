package main

import "strings"

func WordCount(s string) map[string]int {
	store := make(map[string]int)
	words := strings.Fields(s)

	for _, word := range words {
		if value, valueExists := store[word]; valueExists {
			store[word] = value + 1
		} else {
			store[word] = 1
		}
	}
	return store
}
