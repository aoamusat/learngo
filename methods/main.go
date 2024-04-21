package main

import (
	"fmt"
)

func main() {
	fmt.Println("Working with methods and interfaces")
	// Go does not have classes. However, you can define methods on types.

	v := Vertex{float64(1), float64(2)}
	fmt.Printf("Before scaling; Abs value: %.6f\n", v.Abs())

	v.Scale(4.5)
	fmt.Printf("After scaling by factor 3; Abs value: %.6f\n", v.Abs())

	var score CreditScore = 450
	fmt.Printf("Credit score: %d\nScore rating: %s\n\n", score, score.Rating())
}
