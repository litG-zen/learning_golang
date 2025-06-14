package main

import "fmt"

/*
Similar to __str__ method in Python, which is used to define explicit print format for class's object;
We have Stringers(an ubiquitous Interface) in Go, defined by the fmt package.

type Stringer interface {
    String() string
}
*/

// Let's see how we can use it
type Person struct {
	name    string
	age     int
	city    string
	hobbies []string
}

func (p *Person) String() string {
	return fmt.Sprintf("This is %v's profile !\nAge: %v\nCity: %v\nHobbies : %v", p.name, p.age, p.city, p.hobbies)
}

func main() {
	lit_var := Person{name: "Lit", age: 29, city: "Kota", hobbies: []string{"reading", "coding", "Sleeping", "Chess"}}
	fmt.Println(&lit_var)
}
