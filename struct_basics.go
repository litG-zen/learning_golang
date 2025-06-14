package main

import (
	"fmt"
	"math"
)

type Coords struct {
	X int
	Y int
}

func coords_dist(p1, p2 Coords) float64 {
	return math.Sqrt(float64((p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y)))
}

func main() {
	p1 := Coords{1, 1}
	p2 := Coords{5, 4}

	fmt.Printf("Coords P1 details: X = %v, Y = %v\n", p1.X, p1.Y)
	fmt.Printf("Coords P2 details: X = %v, Y = %v\n", p2.X, p2.Y)
	fmt.Printf("Distance between them, %v", coords_dist(p1, p2))

}
