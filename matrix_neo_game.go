package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

const red_pill_logo = `
 __        __   _                            _        
 \ \      / /__| | ___ ___  _ __ ___   ___  | |_ ___  
  \ \ /\ / / _ \ |/ __/ _ \| '_ ' _ \ / _ \ | __/ _ \ 
   \ V  V /  __/ | (_| (_) | | | | | |  __/ | || (_) |
    \_/\_/ \___|_|\___\___/|_| |_| |_|\___|  \__\___/ 
                                                     
              Welcome to Reality, Neo.
`

const blue_pill_logo = `
 ___                                                                                     
  | o ._ _   _    _|_  _     _   _    |_   _.  _ |    _|_  _    o | |      _ o  _  ._  | 
  | | | | | (/_    |_ (_)   (_| (_)   |_) (_| (_ |<    |_ (_)   | | | |_| _> | (_) | | o 
                             _|                                                        
`

func clearScreen() {
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	fmt.Println("Hello Neo! Welcome ... choose one pill option (1:Red, 2:Blue)")
	var pill_opt int
	fmt.Scan(&pill_opt)
	clearScreen()

	switch pill_opt {
	case 1:
		fmt.Printf("Welcome Neo..")
		fmt.Printf("\t You chose Red pill, lets take you out of the world of Matrix.")
		time.Sleep(2)
		clearScreen()
		fmt.Println(red_pill_logo)
	case 2:
		fmt.Printf("Welcome Neo..")
		fmt.Printf("\t You chose Red pill, lets take you out of the world of Matrix.")
		time.Sleep(2)
		clearScreen()
		fmt.Println(blue_pill_logo)

	default:
		defer fmt.Println("Nikal pahli fursat me!")
		fmt.Println("Wrong choice ! You had to choose 1 for Red, or 2 for Blue")

	}

}
