package main

import "fmt"

func fib() func() int {
	var fib_seq []int
	pos := 0
	return func() int {
		if pos == 0 {
			fib_seq = append(fib_seq, 0)
			pos += 1
			return fib_seq[0]

		}
		if pos == 1 {
			fib_seq = append(fib_seq, 1)
			pos += 1
			return fib_seq[1]
		} else {
			fib_seq = append(fib_seq, fib_seq[pos-1]+fib_seq[pos-2])
			pos += 1
			return fib_seq[pos-1]
		}
	}

}

func main() {
	var fib_num int
	fmt.Printf("Enter the number of fibonacci sequences you want!:\t")
	fmt.Scan(&fib_num)
	fib_var := fib()

	for i := 0; i < fib_num; i++ {
		fmt.Printf("\nIndex :%v, Fib number :%v\n", i, fib_var())
	}
}
