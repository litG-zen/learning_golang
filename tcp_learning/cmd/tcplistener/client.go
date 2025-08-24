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
		"GET /connect HTTP1.1 \r\n",
		"Header1 Value1 \r\n",
		"Header2 Value2 \r\n",
		"Header3 Value3 \r\n",
		"Accept */* \r\n",
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
