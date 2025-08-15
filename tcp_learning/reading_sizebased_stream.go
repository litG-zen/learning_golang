package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal("error", "error", err)
	}

	for {
		data := make([]byte, 8)
		n, err := f.Read(data)
		if err != nil {
			fmt.Println("EOF!")
			break
		}
		fmt.Printf("read: %s\n", string(data[:n]))

	}

}
