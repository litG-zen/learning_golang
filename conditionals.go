// Conditionals

// If
// switch

package main

import "fmt"

func main() {

	a := 12 // initiation using walrus operator.

	// Simple if block.
	if a < 12 {
		fmt.Println("Less than 12")
	} else { // please note that else in Go must come after closing bracket(~})/
		fmt.Println("more than 12")
	}

	// Another type of if where initation and check happens at same line
	if a := 13; a < 12 {
		fmt.Println("less than 12")
	} else {
		fmt.Println("more than 12")
	}

	// Switch : using the variable a declared above.
	switch {
	case a < 12:
		fmt.Println("less than 12")
	case a == 12:
		fmt.Println("a is 12")
	default:
		fmt.Println("a is more than 12")
	}
}
