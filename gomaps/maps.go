package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

func main() {
	fmt.Println("Maps/HashMap in Golang")

	m := make(map[string]Vertex)
	m["Natanz Nuclear Facility"] = Vertex{
		40.68433, -74.39967,
	}

	m["Baghram AirField"] = Vertex{
		40.68433, -74.39967,
	}

	m["Nevatim Airbase"] = Vertex{
		40.68433, -74.39967,
	}

	m["Golan Heights"] = Vertex{
		40.68433, -74.39967,
	}

	m["Fordow Nuclear Plant"] = Vertex{
		40.68433, -74.39967,
	}

	m["Google"] = Vertex{
		37.42202, -122.08408,
	}

	for key, location := range m {
		fmt.Printf("%v -> (%f, %f)\n", key, location.Lat, location.Long)
	}

	key := "Google"

	if elem, keyExists := m[key]; keyExists {
		fmt.Printf("%v -> (%f, %f)\n", key, elem.Lat, elem.Long)
	}

	fmt.Printf("%v\n", WordCount("I love You You Love me too"))
}
