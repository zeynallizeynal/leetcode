package main

import "fmt"

type TreeNode struct {
	Value int
	Right *TreeNode
	Left  *TreeNode
}

// aka Depth-First
func RecursiveTraverse(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println("enter ", node.Value)
	RecursiveTraverse(node.Left)
	RecursiveTraverse(node.Right)
	fmt.Println("leave ", node.Value)
}

func PreorderTraverse(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	res := []int{node.Value}
	res = append(res, PreorderTraverse(node.Left)...)
	res = append(res, PreorderTraverse(node.Right)...)
	return res
}
func InorderTraverse(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	res := InorderTraverse(node.Left)
	res = append(res, node.Value)
	res = append(res, InorderTraverse(node.Right)...)
	return res
}
func PostorderTraverse(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	res := PostorderTraverse(node.Left)
	res = append(res, PostorderTraverse(node.Right)...)
	res = append(res, node.Value)
	return res
}

// aka Breadth-First
func LevelOrderTraverse1(root *TreeNode) {
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

func LevelOrderTraverse2(root *TreeNode) {
	if root == nil {
		return
	}

	q := []*TreeNode{root}

	depth := 1

	for len(q) > 0 {
		size := len(q)
		for size > 0 {
			cur := q[0]
			q = q[1:]
			fmt.Printf("depth = %d; val = %d\n", depth, cur.Value)

			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
			size--
		}
		depth++
	}
}

type State struct {
	node  *TreeNode
	depth int
}

func LevelOrderTraverse3(root *TreeNode) {
	if root == nil {
		return
	}

	q := []State{{node: root, depth: 1}}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		fmt.Printf("depth = %d; val = %d\n", cur.depth, cur.node.Value)
		if cur.node.Left != nil {
			q = append(q, State{node: cur.node.Left, depth: cur.depth + 1})
		}
		if cur.node.Right != nil {
			q = append(q, State{node: cur.node.Right, depth: cur.depth + 1})
		}
	}
}
