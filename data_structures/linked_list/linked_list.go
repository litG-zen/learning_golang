package main

import "fmt"

/*
A linked list is a fundamental data structure in computer science. It mainly allows efficient insertion
and deletion operations compared to arrays.

 Like arrays, it is also used to implement other data structures like stack, queue and deque
*/

// We will be creating a simple linked list using Struct and pointers for revision, before we delve into BlockChain.

type LinkedNode struct {
	data int // We will be taking a simple int-value as data for the linked list.
	next *LinkedNode
}

// Functionalities for LinkedList
// Node initiation
// Node addition
// Node search by data
// Node deletion
// Node traversal
// Node reversal

func InitNode(data int) *LinkedNode {
	return &LinkedNode{
		data: data,
		next: nil, // Explicitly setting next to nil (optional, since it's default)
	}
}

func (l *LinkedNode) GetTailNode() *LinkedNode {
	for current := l; current != nil; current = current.next {
		if current.next == nil {
			return current
		}
	}
	return nil
}

func (l *LinkedNode) AddNode(data int) {
	new_node := InitNode(data)
	tail_node := l.GetTailNode()
	tail_node.next = new_node
}

func (l *LinkedNode) SearchData(data int) (bool, *LinkedNode) {

	for current := l; current != nil; current = current.next {
		if current.data == data {
			return true, current
		}
	}
	return false, nil
}

func (l *LinkedNode) Traverse() {
	fmt.Println("\nTraversal Initiated\t")
	for current := l; current != nil; current = current.next {
		fmt.Printf("%v->", current.data)
	}
}

func (l *LinkedNode) DeleteData(data int) {
	present, node := l.SearchData(data)
	if !present {
		fmt.Println("Deletion data not present in the LinkedList")
	} else {
		node_to_delete := node
		for current := l; current != node_to_delete; current = current.next {
			if current.next == node_to_delete {
				fmt.Printf("\n Deleting the node! %v", node_to_delete)
				current.next = node_to_delete.next
				break
			}

		}

	}
}

func (l *LinkedNode) ReverseLinkedList() *LinkedNode {
	// Here it's an assumption that the passed linkedlist is from the header node.

	var prev *LinkedNode
	current := l
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	return prev
}

func main() {
	first_node := InitNode(220396)
	fmt.Printf("First node data :%v", first_node)

	fmt.Println("\n New node addition")
	first_node.AddNode(160896)
	fmt.Printf("\n Updated LinkedList : %v %v\n\n", first_node.data, first_node.next)

	fmt.Println("\n Searching the node for a provided data")
	_, node := first_node.SearchData(160896)
	if node == nil {
		fmt.Println("\n No node with provided data exists! Error !!!!")
	} else {
		fmt.Printf("\n %v has the data", node)
	}

	fmt.Println("\n Searching the node for a non-existing data")
	_, node = first_node.SearchData(12344321)
	if node == nil {
		fmt.Println("\n No node with provided data exists! Error !!!!")
	} else {
		fmt.Printf("\n %v has the data", node)
	}

	for i := 0; i < 10; i++ {
		first_node.AddNode(i * 10)
	}

	first_node.Traverse()

	first_node.DeleteData(160896)

	fmt.Println("\n\n After deletion traversal")
	first_node.Traverse()

	fmt.Println("\n\n Reversing the Linkedlist")
	first_node = first_node.ReverseLinkedList()
	fmt.Println("After Reversal")
	first_node.Traverse()
}
