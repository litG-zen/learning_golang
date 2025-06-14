package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
func word_counter_Panic(str_arr []string) (res map[string]int) {
	// We must initialise the map otherwise it will be nil and will cause panic while running line  16
	for _, word := range str_arr {
		_, ok := res[word]
		if ok == false {
			res[word] = 1
		} else {
			res[word] += 1
		}
	}
	return
}

The above function will not work because the return variable initiation at function signature will fail at line 16
*/

func word_counter(str_arr []string) map[string]int {
	res := make(map[string]int) // We must initialise the map otherwise it will be nil and will cause panic while running line 33
	// var res map[string]int // Uncomment this and comment the above one to cause panic ;)
	for _, word := range str_arr {
		_, present := res[word]
		if !present {
			res[word] = 1
		} else {
			res[word] += 1
		}
	}
	return res
}

func main() {
	var str_arr string

	fmt.Println("Enter the desired sentence to be up for word count !")

	// We cannot use fmt.Scan(&str_arr) here for space-separated string because fmt.Scan() considers space as delimiter.
	// So, it will cause panic issue in the code run when we print the str_arr at line 51
	reader := bufio.NewReader(os.Stdin)
	str_arr, _ = reader.ReadString('\n')

	fmt.Printf("\nEntered sentence\n:%v", str_arr)

	s := strings.Fields(str_arr)
	fmt.Printf("\nWord counts:%v", word_counter(s))
}
