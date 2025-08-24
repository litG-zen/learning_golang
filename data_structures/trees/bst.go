package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func InitTree(data int) *Node {
	return &Node{data, nil, nil}
}

func (root *Node) AddNode(data int) {
	for current_node := root; current_node != nil; {
		if data <= current_node.data {
			if current_node.left == nil {
				current_node.left = &Node{data, nil, nil}
			} else {
				current_node = current_node.left
			}
		}
		if data > current_node.data {
			if current_node.right == nil {
				current_node.right = &Node{data, nil, nil}
			} else {
				current_node = current_node.right
			}
		}
	}
	fmt.Println("addition done")
}

func (root *Node) SearchNode(data int) (bool, *Node) {
	// Search tree in Binary-search mode for the provideed data value.
	// Returns bool and Node, if found.Nil other wise.
	return false, nil
}

func (root *Node) DeleteNode(data int) bool {
	// Search and delete the node.
	return false
}

func (root *Node) InorderTraversal() {
	// Inorder traversal
}

func (root *Node) PreorderTraversal() {
	// Prerder traversal
}

func (root *Node) PostorderTraversal() {
	// Postorder traversal
}

func main() {
	root := InitTree(22)
	root.AddNode(16)

	fmt.Println(root)
}
