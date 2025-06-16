package main

import (
	"fmt"
)

// This is an attempt to solve Celebrity problem(https://www.geeksforgeeks.org/dsa/the-celebrity-problem/) using Stacks in Go
// Time Complexity : O(n) and Space Complexity : O(n)

// defining methods for Stack struct
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) IsFull() bool {
	return len(s.items) == cap(s.items)
}

func (s *Stack) Peek() (int, bool) {
	if s.IsEmpty() {
		return -1, true
	}
	return s.items[len(s.items)-1], false
}

func (s *Stack) Push(val int) {
	if s.IsFull() {
		fmt.Println("Stack OverFlow!")
	}
	s.items = append(s.items, val)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}
	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val
}

type Stack struct {
	items []int
}

func NewStackCreate(size int) *Stack {
	return &Stack{
		items: make([]int, 0, size),
	}
}

func SelectCelebrity(mat [][]int) int {
	// returns index of row (denoting that person is celebrity); returns -1 if no celeb found
	stack := NewStackCreate(len(mat[0]))

	for i := 0; i < len(mat[0]); i++ {
		stack.items = append(stack.items, i)
	}

	for len(stack.items) > 1 {
		a := stack.Pop()
		b := stack.Pop()

		if mat[a][b] != 0 {
			stack.Push(b)
		} else {
			stack.Push(a)
		}
	}
	c := stack.Pop()
	for i := 0; i < len(mat[0]); i++ {
		if i != c {
			if mat[c][i] == 1 || mat[i][c] == 0 {
				return -1
			}
		}
	}
	return c

}
func main() {
	// Initialising a 2D Square matrix
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
