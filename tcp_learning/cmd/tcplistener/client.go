// tcp_client.go
package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Send some messages
	messages := []string{
		"KYA bolti public\n",
		"Machayenge naaa\n",
		"Wow, streaming data is fun!\n",
	}

	for _, msg := range messages {
		fmt.Printf("sending: %s", msg)
		_, err := conn.Write([]byte(msg))
		if err != nil {
			panic(err)
		}
		time.Sleep(1 * time.Second) // slow down messages
	}
}
