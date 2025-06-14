package main

import (
	"fmt"
	"os"
)

const MAX_ARR_SIZE = 20

func main() {
	var ARR_SIZE int
	fmt.Printf("Enter desired array size(MaxAllowedSize : %v)\t", MAX_ARR_SIZE)
	fmt.Scan(&ARR_SIZE)
	if ARR_SIZE > MAX_ARR_SIZE {
		fmt.Println("passed desired array size is higher than the Max allowed array size")
		os.Exit(1)
	}

	// make([]type, size) can be used to dymamically allocate memory array
	/*

		In comparison to C's malloc

		| C                                     | Go                      |
		| ------------------------------------- | ----------------------- |
		| `int *p = malloc(sizeof(int));`       | `p := new(int)`         |
		| `int *arr = malloc(n * sizeof(int));` | `arr := make([]int, n)` |
		| `free(p)`                             | No need (GC handles it) |

	*/
	a := make([]int, ARR_SIZE)

	for i := 0; i < ARR_SIZE; i++ {
		fmt.Printf("Enter the %vth index value", i)
		fmt.Scan(&a[i])
	}

	fmt.Printf("The Final array %v", a)

	var slice_choice, slice_index int

	fmt.Printf("\nDo you want to attempt array_slicing? 1/0\t")
	fmt.Scan(&slice_choice)
	if slice_choice == 0 {
		fmt.Printf("\nCool!, we won't slice ")
	} else {
		fmt.Printf("\nEnter the slice_index (starting from 0)")
		fmt.Scan(&slice_index)
		if slice_index > ARR_SIZE-1 {
			fmt.Printf("Slicing can't happen  bro! this is not Python!")
			os.Exit(1)
		} else {
			fmt.Printf("\nSliced array %v", a[:slice_index])

			// We have simply used Print statement to slice, mentioned below is the approach to initate a slice array
			var s []int = a[:slice_index]
			fmt.Printf("\nInitiated slice array: %v", s)

		}
	}

}
