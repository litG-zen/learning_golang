package main
import (
	"fmt"
	"os"
	"syscall"
)


func main(){	

	var connected_clients int = 0

	var max_client int = 20000

	var events []syscall.EpollEvent = make([]syscall.EpollEvent, max_client)

	serverFD, err := syscall.Socket(syscall.AF_INET, syscall.O_NONBLOC | syscall.SOCK_STREAM, 0)
	if err != nil {
		fmt.Printf("socket error: %s\n", err)
		os.Exit(1)
	}

	syscall.SetNonblock(serverFD, true)	
	
}
