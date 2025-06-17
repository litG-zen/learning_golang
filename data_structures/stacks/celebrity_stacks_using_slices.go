package main

import "fmt"

func SelectCelebrity(mat [][]int) int {
	stack := make([]int, 0, len(mat[0]))
	for i := 0; i < len(mat[0]); i++ {
		stack = append(stack, i)
	}

	for len(stack) > 1 {
		a := stack[len(stack)-1]
		b := stack[len(stack)-2]

		stack = stack[:len(stack)-2]
		if mat[a][b] == 1 {
			stack = append(stack, b)
		} else {
			stack = append(stack, a)
		}
	}
	c := stack[0]
	for i := 0; i < len(mat[0]); i++ {
		if i == c {
			continue
		}
		if mat[c][i] == 1 || mat[i][c] == 0 {
			return -1
		}
	}
	return c
}

func main() {
	mat_1 := [][]int{
		{1, 1, 1},
		{0, 1, 0},
		{1, 1, 1},
	}
	fmt.Printf("\nMatrix : %v", mat_1)
	celeb := SelectCelebrity(mat_1)
	if celeb != -1 {
		fmt.Printf("\n\nCelebrity :%v", celeb)
	} else {
		fmt.Println("No celebrity found!")
	}
}
