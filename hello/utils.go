package main

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

func suffixValue(length int) string {
	rand1 := rand.Intn(1000000000)
	rand2 := int(math.Pow10(int(math.Abs(float64(length/2-1))))) +
		rand.Intn(9*int(math.Pow10(int(math.Abs(float64(length/2-1))))))
	rdx := strconv.Itoa(rand1 * rand2)
	return rdx[:length]
}

func generateSessionID() string {
	prefix := "010345"
	if prefix == "" {
		prefix = suffixValue(6)
	}

	currTime := time.Now().Format("060102150405") // YYMMDDHHmmss format
	sfx := suffixValue(12)
	return prefix + currTime + sfx
}
