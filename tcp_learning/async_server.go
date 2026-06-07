package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"syscall"
)

const (
	HOST = "127.0.0.1"
	PORT = 8000
)

// //////////////////////////////////////////////////////////////////////////////////

type FDComm struct {
	Fd int
}

func (c *FDComm) Read(b []byte) (int, error) {
	return syscall.Read(c.Fd, b)
}

func (c *FDComm) Write(b []byte) (int, error) {
	return syscall.Write(c.Fd, b)
}

// //////////////////////////////////////////////////////////////////////////////////

// Socket read and write functions
func readCommand(c io.ReadWriter) (string, error) {
	var buf []byte = make([]byte, 1024)
	n, err := c.Read(buf)
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}

func writeCommand(c io.ReadWriter, response string) error {
	_, err := c.Write([]byte(response))
	return err
}

func RunAsyncTCPServer() error {
	var connected_clients int = 0

	var max_client int = 20000

	// Epoll event array. This buffer will be holdind FD that are ready for i/o operations by system call.
	var events []syscall.EpollEvent = make([]syscall.EpollEvent, max_client)

	// Create a TCP socket.

	// AF_INET: IPv4 protocol
	// SOCK_STREAM: This tells the system that we want to keep the connection open and send data in a stream format.
	// O_NONBLOCK: This flag makes the socket non-blocking,
	// which means that system calls on this socket will return immediately if they cannot be completed,
	// rather than blocking the execution of the program until they can be completed.

	serverFD, err := syscall.Socket(syscall.AF_INET, syscall.O_NONBLOCK|syscall.SOCK_STREAM, 0)
	if err != nil {
		fmt.Printf("socket error: %s\n", err)
		os.Exit(1)
	}
	defer syscall.Close(serverFD) // close the socket when the connection is closed.

	// Set the socket to non-blocking mode.
	// This allows the server to handle multiple clients concurrently without blocking on any single client.
	if err = syscall.SetNonblock(serverFD, true); err != nil {
		return err
	}

	// Bind the IP and port to the socket.
	ip := net.ParseIP(HOST)
	fmt.Printf("Server is listening on %s:%d:%s\n", HOST, PORT, ip)
	if err = syscall.Bind(serverFD, &syscall.SockaddrInet4{
		Port: PORT,
		Addr: [4]byte{ip[0], ip[1], ip[2], ip[3]},
	}); err != nil {
		return err
	}

	// Listen for incoming connections.
	if err = syscall.Listen(serverFD, max_client); err != nil {
		return err
	}

	// Asyncio server loop begins here.
	epollFD, err := syscall.EpollCreate1(0)
	if err != nil {
		fmt.Printf("epoll create error: %s\n", err)
	}
	defer syscall.Close(epollFD)

	// We now want the server(serverFD) to be monitored by epoll for incoming connections.
	var socketServerEvent syscall.EpollEvent = syscall.EpollEvent{
		Events: syscall.EPOLLIN, // EPOLLIN: This event is triggered when there is data to read on the socket.
		Fd:     int32(serverFD), // The file descriptor of the server socket that we want to monitor for incoming connections.
	}

	// We add the server socket to the epoll instance using the EpollCtl function.
	if err = syscall.EpollCtl(epollFD, syscall.EPOLL_CTL_ADD, serverFD, &socketServerEvent); err != nil {
		return err
	}

	for {
		// Wait for events on the monitored file descriptors. This call will block until at least one event occurs.
		eventCount, err := syscall.EpollWait(epollFD, events, -1)
		if err != nil {
			return err
		} else {
			for i := 0; i < eventCount; i++ {
				event := events[i]
				if event.Fd == int32(serverFD) {
					// This means that there is a new incoming connection on the server socket.
					connFD, client_addr, err := syscall.Accept(serverFD)
					if err != nil {
						fmt.Printf("accept error: %s\n", err)
						continue
					}

					connected_clients++
					fmt.Printf("New client connecte. Total clients: %d\n", connected_clients)

					switch addr := client_addr.(type) {
					case *syscall.SockaddrInet4:
						fmt.Printf("New connection from %s:%d\n", net.IP(addr.Addr[:]), addr.Port)
					case *syscall.SockaddrInet6:
						fmt.Printf("New connection from %s:%d\n", net.IP(addr.Addr[:]), addr.Port)
					default:
						fmt.Printf("Unknown client address type\n")
					}
					syscall.SetNonblock(serverFD, true)

					var clientEvent syscall.EpollEvent = syscall.EpollEvent{
						Events: syscall.EPOLLIN, // EPOLLIN: This event is triggered when there is data to read on the socket.
						Fd:     int32(connFD),   // The file descriptor of the client socket that we want to monitor for incoming data.
					}
					// register the client socket with epoll to monitor for incoming data.
					if err = syscall.EpollCtl(epollFD, syscall.EPOLL_CTL_ADD, connFD, &clientEvent); err != nil {
						fmt.Printf("epoll ctl error: %s\n", err)
						continue
					}
				} else {
					comm := FDComm{Fd: int(event.Fd)}
					cmd, err := readCommand(&comm)
					if err != nil {
						fmt.Printf("read error: %s\n", err)
						syscall.Close(int(event.Fd))
						connected_clients--
						fmt.Printf("Client disconnected. Total clients: %d\n", connected_clients)
						continue
					}
					writeCommand(&comm, cmd)
				}

			}
		}
	}
}

func main() {
	err := RunAsyncTCPServer()
	if err != nil {
		fmt.Printf("Server error: %s\n", err)
	}
}
