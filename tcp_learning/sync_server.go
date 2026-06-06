package main

import (
	"fmt"
	"net"
)

func main() {
	listner, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Listening on port 8000")

	for {
		var buff []byte = make([]byte, 512)
		conn, err := listner.Accept()
		if err != nil {
			fmt.Printf("line 20 err", err)
		}
		fmt.Printf("connection established with %s\n", conn.RemoteAddr())
		for {
			n, err := conn.Read(buff[:])
			if err != nil {
				fmt.Printf("line 24 err", err)
			}
			command := string(buff[:n])
			if _, err := conn.Write([]byte(command)); err != nil {
				fmt.Printf("line 29 err", err)
			}
		}
	}
}