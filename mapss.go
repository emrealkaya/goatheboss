package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{40.68433, -74.39967},
	"Google":    Vertex{37.42202, -122.08408},
}

var key = map[string]int{
	"one": 1,
	"two": 2,
}

func main() {

	key["one"] = 10
	key["four"] = 4
	key["three"] = 3
	fmt.Println(key)

	key, ok := key["two"]
	fmt.Println(key, ok)

	// m = make(map[string]Vertex)
	// m["Bell Labs"] = Vertex{40.68433, -74.39967}
	// fmt.Println(m["Bell Labs"])
	fmt.Println(m)
}
