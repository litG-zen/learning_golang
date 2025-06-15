package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("This is a great language!")

	b := make([]byte, 5)

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v, err = %v bytes_read = %v", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
