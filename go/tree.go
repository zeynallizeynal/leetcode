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

func LevelOrderTraverse(root *TreeNode) {
	if root == nil {
		return
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		fmt.Println(cur.Value)

		if cur.Left != nil {
			q = append(q, cur.Left)
		}
		if cur.Right != nil {
			q = append(q, cur.Right)
		}
	}
}
