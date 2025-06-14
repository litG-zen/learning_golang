/*
// Function closures

A closure is a function value that references variables from outside its body.
The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
*/

package main

import "fmt"

func counter() func() int {
	counter := 0
	return func() int {
		counter += 1
		return counter
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	counter_var := counter()
	fmt.Println(counter_var())
	fmt.Println(counter_var())
	fmt.Println(counter_var())
	fmt.Println(counter_var())

	val, exp := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("\nVal :%v, exp :%v", val(i), exp(i*i))
	}
}
