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
	post_body := "{'body':'this is a text', 'test':'this is a test'}"
	// Send some messages
	messages := []string{
		/*
			As per RFC 9110 semantics, the HTTP connection request is nothing but a steam of data with \r\n
			as the separator.
			The first line represents API method, path and HTTP version used definition.
			The later lines represent headers.
			A single line where the line content is only \r\n represents request separator
			than means that the lines that follow will contain request body ( if any ).
		*/
		"POST /send_data HTTP1.1 \r\n",
		"Content-Type application/json \r\n",
		"Content-Length" + string(len(post_body)) + "\r\n",
		"Header3 Value3 \r\n",
		"Accept */* \r\n",
		"\r\n",
		post_body,
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
