// reading the data from a file over TCP

package main

import (
	"fmt"
	"log"
	"net"
	"tcp_learning/cmd/internal/requests"
)

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
		r, err := requests.RequestFromReader(conn)
		if err != nil {
			log.Fatal("Request failed", "Error", err)
		}
		fmt.Printf("Request Line: \n")
		fmt.Printf("- Method: %s\n", r.RequestLine.Method)
		fmt.Printf("- Target: %s\n", r.RequestLine.RequestTarget)
		fmt.Printf("- Version: %s\n", r.RequestLine.HttpVersion)

	}

}

// to run and test this script, do
// Shell 1: go run main_requests.go
// Shell 2: curl http://localhost:8080/health
