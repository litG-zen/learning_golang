package main

import (
	"fmt"
	"time"
)

// Error handling in Go
/*
Similar to Stinger interface, Go also provides an error interface.
*/

type MyError struct {
	When time.Time
	What string
}

func (m *MyError) Error() string {
	return fmt.Sprintf("at %v, %v", m.When, m.What)
}

func run() error {
	return &MyError{time.Now(), "this is a error !"}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
