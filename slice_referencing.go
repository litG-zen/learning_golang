/*
This program(not Script; because script term is used for interpreted languages like Python, Bash, JS)
is to demonstrate how Golang slicing uses array reference in place of creating a completly new array.
*/

package main

import "fmt"

func main() {

	// Array initiatiion
	arr := [6]int{12, 24, 36, 48, 60, 72}

	fmt.Printf("Initated array: \t%v", arr)

	slice_1 := arr[0:4] // We can also initiate by slice_1 []int= arr[0:3], here we are explicitly defining  type, but Golang supports both the ways
	slice_2 := arr[2:5]

	fmt.Printf("\nSlice 1: \t%v", slice_1)
	fmt.Printf("\nSlice 2: \t%v", slice_2)

	fmt.Printf("\n")

	// Attempting value update in 3rd index of slice_1(3rd index of original array and 1st index of slice_2 will get updated)
	slice_1[3] = 22
	fmt.Println(slice_1[3], arr[3], slice_2[1])

	/*
		A slice has both a length and a capacity.

		The length of a slice is the number of elements it contains.

		The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

		The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).

		You can extend a slice's length by re-slicing it, provided it has sufficient capacity.
	*/

	fmt.Printf("\nSlice_1's details\n Slice data :%v\n Slice len :%v,\n Slice capacity :%v",
		slice_1, len(slice_1), cap(slice_1))

	// Reshaping Slice_1
	slice_1 = arr[:]
	fmt.Printf("\nPost Reshape Slice_1's details\n Slice data :%v\n Slice len :%v,\n Slice capacity :%v",
		slice_1, len(slice_1), cap(slice_1))

}
