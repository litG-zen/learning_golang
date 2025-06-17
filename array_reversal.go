package main

import "fmt"

func main() {

	// Array reversal

	str := "hello world"
	arr := []byte(str)

	fmt.Printf("Original bytearray %q\n", arr)
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}

	fmt.Printf("\nReversed bytearray %q\n", arr)
}
