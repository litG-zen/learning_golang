package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal("error", "error", err)
	}
	str := ""
	for {
		data := make([]byte, 8)
		n, err := f.Read(data)
		if err != nil {
			fmt.Println("EOF!")
			break
		}
		// Reading splitting at new-line.
		data = data[:n]
		if i := bytes.IndexByte(data, '\n'); i != -1 {
			str += string(data[:n])
			data = data[i+1:]
			fmt.Printf("read %s\n", str)
			str = ""
		}
		str += string(data)
	}
	if len(str) != 0 {
		fmt.Printf("read %s\n", str)

	}
}
