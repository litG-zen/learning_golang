package main

import (
	"fmt"
	"math/rand"
)

// Stack using Arrays

// This program is a failed version of stack using Go by using slices, because the slice intiation to STACK_SIZE
// will create an array of 0s, resulting in StackOverflow when we attempt Push.

/*
Basic Operations:
Push: Adds a new element to the top of the stack.
Pop: Removes the top element from the stack.
Peek/Top: Returns the value of the top element without removing it.
isEmpty: Checks if the stack is empty.
isFull: Checks if the stack is full (relevant when using arrays).
*/

const STACK_SIZE = 10

func main() {
	stack := make([]int, STACK_SIZE)

	fmt.Printf("Stack initaited: data:%v capacity:%v current_size :%v; ", stack, cap(stack), len(stack))

	fmt.Printf("\nStack filing initiated!")
	for i := 0; i < 5; i++ {
		fmt.Printf("\n Iter:%v", i+1)
		Push(rand.Intn(STACK_SIZE*10), stack)
	}

	fmt.Printf("Stack status after filling: data:%v capacity:%v current_size :%v; ", stack, cap(stack), len(stack))

}

func IsEmpty(stack []int) bool {
	if len(stack) == 0 {
		return true
	}
	return false
}

func IsFull(stack []int) bool {
	if len(stack) == STACK_SIZE {
		return true
	}
	return false
}

func Push(val int, stack []int) []int {
	temp_stack := stack
	if !IsFull(stack) {
		temp_stack = append(stack, val)
	} else {
		fmt.Println("Stack Overflow!")
	}
	return temp_stack
}

func Pop(stack []int) int {
	if IsEmpty(stack) {
		fmt.Println("Stack in empty!")
		return -1
	} else {
		val := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return val
	}
}
