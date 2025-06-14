package main

import (
	"fmt"
	"math"
)

// Interface: A set of defined method signatures. It is in a way, Go's way of polymorphism without OOPs.
// Area() and Perimeter() can be two methods that can act diff for Circle and Rectangle struct.

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Len, Width float64
}

type Square struct {
	Len float64
}

func (c *Circle) Area() float64 {
	if c == nil { //null pointer handling
		fmt.Println("<NIL>")
		return 0.0
	}
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	if c == nil { //null pointer handling
		fmt.Println("<NIL>")
		return 0.0
	}
	return 2 * math.Pi * c.Radius
}

func (r *Rectangle) Area() float64 {
	if r == nil { //null pointer handling
		fmt.Println("<NIL>")
		return 0.0
	}
	return r.Len * r.Width
}

func (r *Rectangle) Perimeter() float64 {
	if r == nil { //null pointer handling
		fmt.Println("<NIL>")
		return 0.0
	}
	return 2 * (r.Len + r.Width)
}

func (s *Square) Area() float64 {
	if s == nil { //null pointer handling
		fmt.Println("<NIL>")
		return 0.0
	}
	return s.Len * s.Len
}

func (s *Square) Perimeter() float64 {
	if s == nil { //null pointer handling
		fmt.Println("<NIL>")
		return 0.0
	}
	return 4 * s.Len
}

func main() {
	var s Shape

	s = &Circle{Radius: 10} // Circle implements shape
	/*
		Note
		s = Circle{Radius:10} type initialization will fail to initialise
		We are assigning a value type Circle to s, but the Circle type implements the Shape interface using pointer receivers:
		func (c *Circle) Area() float64
		func (c *Circle) Perimeter() float64
		So Go expects a *Circle, not a Circle.
	*/
	fmt.Printf("\nCircle features: \tArea:%v \tPerimeter:%v", s.Area(), s.Perimeter())

	s = &Square{Len: 5}
	fmt.Printf("\nSquare features: \tArea:%v \tPerimeter:%v", s.Area(), s.Perimeter())

	s = &Rectangle{Len: 5, Width: 10}
	fmt.Printf("\nRectangle features: \tArea:%v \tPerimeter:%v", s.Area(), s.Perimeter())

	// Testing null-pointer exception in interface functions
	fmt.Printf("\n")
	var c_nil *Circle
	s = c_nil
	fmt.Printf("\nNilCircle features: \tArea:%v \tPerimeter:%v", s.Area(), s.Perimeter())

}
