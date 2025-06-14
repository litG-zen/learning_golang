package main

import "fmt"

func main() {
	a := [10]int{0, 2, 4, 8, 16, 32, 64, 128, 256, 512}

	// Two ways of iteration

	// Using len(a); similar as in Python
	for i := 0; i < len(a); i++ {
		fmt.Printf("%v ", a[i])
	}

	// Using range form of the for loop, which is similar to ennumarate(arr) in Python , which returns index,value
	for i, v := range a {
		fmt.Printf("\nIndex :%v, Value :%v ", i, v)
	}

}
