package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Stack using Arrays

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
	stack := make([]int, 0, STACK_SIZE)
	var popped_val int
	fmt.Printf("Stack initaited: data:%v capacity:%v current_size :%v; ", stack, cap(stack), len(stack))

	top_val, is_empty := Peek(stack)

	fmt.Printf("\nStack's topmost value! value: %v, is_stack_empty: %v", top_val, is_empty)

	fmt.Printf("\nStack filing initiated!")
	for i := 0; i < 5; i++ {
		fmt.Printf("\n Iter:%v", i+1)
		stack = Push(rand.Intn(STACK_SIZE*10), stack)
	}

	fmt.Printf("\nStack status after filling: data:%v capacity:%v current_size :%v; ", stack, cap(stack), len(stack))

	top_val, is_empty = Peek(stack)
	fmt.Printf("\nStack's topmost value! value: %v, is_stack_empty: %v", top_val, is_empty)

	time.Sleep(500 * time.Microsecond)
	fmt.Printf("\nInitiating stack popping")
	for {
		popped_val, stack = Pop(stack)
		if popped_val == -1 {
			break
		}
		fmt.Printf("\nPopped value : %v", popped_val)
	}

}

func Peek(stack []int) (int, bool) {
	if IsEmpty(stack) {
		return -1, true
	}
	return stack[len(stack)-1], false
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

func Pop(stack []int) (int, []int) {
	if IsEmpty(stack) {
		fmt.Println("Stack is empty!")
		return -1, stack
	} else {
		val := stack[len(stack)-1]
		stack = stack[:(len(stack) - 1)]
		return val, stack
	}
}
