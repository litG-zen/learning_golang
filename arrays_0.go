package main

import "fmt"

func sort_pointer(a *[6]int, reverse bool) {
	var temp int
	var size = 6
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if (a[i] > a[j]) && reverse {
				temp = a[j]
				a[j] = a[i]
				a[i] = temp
			}
			if (a[i] < a[j]) && !reverse {
				temp = a[j]
				a[j] = a[i]
				a[i] = temp
			}
		}
	}
}

func main() {
	a := [6]int{12, 4, 32, 6, 1, 22}
	fmt.Println(a)
	sort_pointer(&a, true)
	fmt.Printf("Asc sort %v", a)
	sort_pointer(&a, false)
	fmt.Printf("\nDesc sort %v", a)
}
