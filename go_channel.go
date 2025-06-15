package main

import "fmt"

/*
Go Channels are typed pipes, acting as medium of passing data.

ch <- v : this  is used to pass value of v to channel
v := <-c : this is used to read from the value of v

Note: Like maps and slices(two data structures of Go), channels must also be created before  use
*/

func sum(s []int, c chan int) {
	sum := 0
	for _, i := range s {
		sum += i
	}
	c <- sum
}

func main() {
	c := make(chan int)
	s := []int{1, -4, 5, 10, 13, 14}

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
