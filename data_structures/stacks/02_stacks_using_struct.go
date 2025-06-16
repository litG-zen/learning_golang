package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Stacks using struct

const STACK_SIZE = 10

type Stack struct {
	items []int
}

func NewStack() *Stack {
	return &Stack{
		items: make([]int, 0, STACK_SIZE),
	}
}

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

func (stack *Stack) ReverseStack() *Stack {
	// This function takes a stack and returns its reverse.
	// i.e [1,2,3,4,5](5 being topmost) becomes [5,4,3,2,1](1 being topmost)

	temp_stack := NewStack()

	for !(stack.IsEmpty()) {
		temp_stack.items = append(temp_stack.items, stack.Pop())
	}
	return temp_stack
}
func main() {
	// initialising an empty stack with STACK_SIZE capacity
	stack := NewStack()

	stack.Push(rand.Intn(100)) // Inserting random number between 0-100
	stack.Push(rand.Intn(100)) // Inserting random number between 0-100
	stack.Push(rand.Intn(100)) // Inserting random number between 0-100
	stack.Push(rand.Intn(100)) // Inserting random number between 0-100
	stack.Push(rand.Intn(100)) // Inserting random number between 0-100

	fmt.Printf("\nStack stats :\n Data: %v\n Size: %v\n Capacity:%v", stack.items, len(stack.items), cap(stack.items))

	fmt.Printf("\nReversing the stack!")
	fmt.Printf("\nStack before reversal: %v", stack.items)
	stack = stack.ReverseStack()
	fmt.Printf("\nStack after reversal: %v", stack.items)

	time.Sleep(100 * time.Millisecond)
	fmt.Printf("\n\nPopping items!\n")
	for !stack.IsEmpty() {
		fmt.Println(stack.Pop())
	}

}
