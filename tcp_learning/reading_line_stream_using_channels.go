package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string, 1)

	go func() {
		defer f.Close()
		defer close(out)
		str := ""
		for {
			data := make([]byte, 8)
			n, err := f.Read(data)
			if err != nil {
				break
			}
			// Reading splitting at new-line.
			data = data[:n]
			if i := bytes.IndexByte(data, '\n'); i != -1 {
				str += string(data[:n])
				data = data[i+1:]
				out <- str
				str = ""
			}
			str += string(data)
		}
		if len(str) != 0 {
			out <- str
		}

	}()
	return out
}

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal("error", "error", err)
	}
	lines := getLinesChannel(f)
	for line := range lines {
		fmt.Println(line)
	}
}
