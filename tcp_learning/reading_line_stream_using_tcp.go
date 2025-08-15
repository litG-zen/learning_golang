// reading the data from a file over TCP

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
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
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("error: ", err)

	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("listening :", err)

		}
		for line := range getLinesChannel(conn) {
			fmt.Printf("read : %s\n", line)
		}
	}

}
