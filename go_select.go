/*
The select statement lets a goroutine wait on multiple communication operations.
A select blocks until one of its cases can run, then it executes that case.
It chooses one at random if multiple are ready.
*/

package main

import (
	"fmt"
	"time"
)

func slowTask(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "Finished slow task"
}

func fastTask(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Finished fast task"
}

func main() {
	slowChan := make(chan string)
	fastChan := make(chan string)

	go slowTask(slowChan)
	go fastTask(fastChan)

	select {
	case msg := <-slowChan:
		fmt.Println("slowChan sent:", msg)
	case msg := <-fastChan:
		fmt.Println("fastChan sent:", msg)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout!")
	}
}
