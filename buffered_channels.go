package main

import "fmt"

type Cars struct {
	company string
	price   float64
	color   string
}

func main() {
	// Buffered channels
	ch := make(chan Cars, 3)

	ch <- Cars{company: "Tesla", price: 45000, color: "red"}
	ch <- Cars{company: "BMW", price: 54000, color: "black"}
	ch <- Cars{company: "Porche", price: 100000, color: "yellow"}

	for i := 0; i < 3; i++ {
		fmt.Printf("Channel %v value: %v\n", i+1, <-ch)

	}

}
