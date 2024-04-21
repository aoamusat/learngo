package main

import "math"

type CreditScore int16

func (score CreditScore) Rating() string {
	if score >= 300 && score <= 499 {
		return "Poor"
	} else if score >= 500 && score <= 649 {
		return "Average"
	} else if score >= 650 && score <= 749 {
		return "Good"
	} else if score >= 750 && score <= 850 {
		return "Excellent"
	} else {
		return "Invalid score"
	}
}

type Vertex struct {
	X float64
	Y float64
}

// Defines method Abs() on the Vertex struct
func (v Vertex) Abs() float64 {
	var AbsValue float64 = math.Sqrt(v.X*v.X + v.Y*v.Y)
	return AbsValue
}

func (v Vertex) EuclideanDistance() float64 {
	return math.Sqrt(v.Y*v.Y - v.X*v.X)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
