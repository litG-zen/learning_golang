/*
Similar to Python, we can also append data in a slice in Golang. This program is a demo of that.
*/

package main

import "fmt"

func print_slice_details(a []int) {
	fmt.Printf("\nSlice data :%v\nSlice len :%v\nSlice capacity :%v\n\n", a, len(a), cap(a))
}

func main() {
	var a []int
	print_slice_details(a)

	// Appending to slice
	a = append(a, 1)
	print_slice_details(a)

	// Appending multiple elements to the slice
	a = append(a, 2, 3, 4, 5)
	print_slice_details(a)

	// Can we append a slice into another slice
	// a = append(a, a[:4]) // No this won't work
}
