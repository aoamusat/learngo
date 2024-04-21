package main

func modify(x *int) int {
	*x += 1
	return *x
}
