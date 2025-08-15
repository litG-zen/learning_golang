package main

import (
	"fmt"
	"time"
)

func sendChan(ch chan string) {
	fmt.Println("Starting sendChan")
	ch <- "hello"
	ch <- "World"
	ch <- "this"
	ch <- "is"
	ch <- "Go"
	ch <- "Channel"
	fmt.Println("\nSendChan finished")
}

func getChan(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Printf("\n %v", input)
	}

}
func main() {

	ch := make(chan string)
	go sendChan(ch)
	go getChan(ch)
	time.Sleep(1 * time.Second)
}
