package main

import "fmt"

type fourwheeler struct {
	mileage       float64
	tank_capacity int
	is_electric   bool
}

type car struct {
	fourwheeler
	company string
	cost    float64
	color   string
}

func main() {
	tesla := car{
		fourwheeler: fourwheeler{
			mileage:       45,
			tank_capacity: 50,
			is_electric:   true,
		},
		company: "Tesla",
		cost:    450000,
		color:   "red",
	}

	fmt.Println(tesla)
}
