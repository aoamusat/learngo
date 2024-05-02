package main

import (
	"fmt"
	"time"
)

func main() {
	// go spinner(100 * time.Millisecond)
	// const n = 45
	// fibN := fib(n) // slow
	// fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
	f("Direct")
	go f("Goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
