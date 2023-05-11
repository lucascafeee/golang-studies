package main

import (
	"fmt"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

func NewNode(data int) *Node {
	return &Node{
		data:  data,
		left:  nil,
		right: nil,
	}
}

func PreorderTraversal(root *Node) {
	if root != nil {
		fmt.Printf("%d ", root.data)
		PreorderTraversal(root.left)
		PreorderTraversal(root.right)
	}
}

func InorderTraversal(root *Node) {
	if root != nil {
		InorderTraversal(root.left)
		fmt.Printf("%d ", root.data)
		InorderTraversal(root.right)
	}
}

func PostorderTraversal(root *Node) {
	if root != nil {
		PostorderTraversal(root.left)
		PostorderTraversal(root.right)
		fmt.Printf("%d ", root.data)
	}
}

func main() {
	root := NewNode(1)
	root.left = NewNode(2)
	root.right = NewNode(3)
	root.left.left = NewNode(4)
	root.left.right = NewNode(5)

	fmt.Println("Preorder Traversal:")
	PreorderTraversal(root)
	fmt.Println()

	fmt.Println("Inorder Traversal:")
	InorderTraversal(root)
	fmt.Println()

	fmt.Println("Postorder Traversal:")
	PostorderTraversal(root)
	fmt.Println()
}
