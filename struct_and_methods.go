/*
Since Go doesn't have the concept of Classes, the method for a struct must be  defined separately.
*/

package main

import "fmt"

type Person struct {
	name   string
	age    int
	city   string
	gender string
}

/*
Note here that, method declaration is similar to a function declaration, but the only diff being that
the method name is defined with a receiver.

For example, we define a method greetMethod and a function greetFunc; both of them perform the same job, but see how they are defined
*/

// method declaration
func (p *Person) greetMethod() string {
	return "Hey!, " + p.name
}

func greetFunc(p *Person) string {
	return "Hey!, " + p.name
}

func main() {
	lalit_var := Person{"Lalit Gupta", 29, "Kota", "Male"}

	fmt.Println(lalit_var.greetMethod())

	// One thing to notice here is that a function defined with Pointer must take a pointer
	fmt.Println(greetFunc(&lalit_var))

	// But a method defined with Pointer can work for both value and a pointer, the line 35 call is value based.
	// Pointer based method invocation
	p := &lalit_var
	fmt.Println(p.greetMethod())

}
