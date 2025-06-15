package main

import "fmt"

/*
Just like templates in C++, where any datatype can be handled within a single function, Go has type-parameters.
*/

func isPresent[T comparable](s []T, x T) bool {
	for _, v := range s {
		if v == x {
			return true
		}
	}
	return false
}

func main() {
	// Trying for int
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 12}
	fmt.Println(isPresent(s, 15))

	// Trying for string
	i := []string{"this", "is", "a", "string"}
	fmt.Println(isPresent(i, "string"))

}
