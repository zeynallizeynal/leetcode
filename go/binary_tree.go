package main

import "fmt"

type TreeNode struct {
	Value int
	Right *TreeNode
	Left  *TreeNode
}

func Traverse(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println("enter ", node.Value)
	Traverse(node.Left)
	Traverse(node.Right)
	fmt.Println("leave ", node.Value)
}
