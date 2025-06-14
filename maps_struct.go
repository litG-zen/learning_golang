package main

import (
	"fmt"
	"math"
)

type Coords struct {
	X, Y int
}

func coords_dist(p1, p2 Coords) float64 {
	return math.Sqrt(float64((p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y)))
}

func main() {
	var coords = map[string]Coords{
		"coords1": {1, 2},
		"coords2": {5, 5},
	}

	fmt.Println(coords["coords1"])
	fmt.Println(coords["coords2"])

	fmt.Println(coords_dist(coords["coords1"], coords["coords2"]))

	delete(coords, "coords1")
	fmt.Println(coords["coords1"])

	v, ok := coords["coords1"]
	fmt.Println(v, ok)

	v_1, ok_1 := coords["coords2"]
	fmt.Println(v_1, ok_1)

}
