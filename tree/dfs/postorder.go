package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	root := buildTreeFromSlice([]any{5, 3, 6, 2, 4, nil, 8, 1, nil, nil, nil, 7, 9})
	fmt.Println(postorder(root))
	fmt.Println(iterative(root))
	fmt.Println(inorderTraversal(root))
}

func postorder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	resp := []int{}

	left := postorder(root.Left)
	right := postorder(root.Right)
	resp = append(resp, left...)
	resp = append(resp, right...)
	resp = append(resp, root.Val)

	return resp
}

func iterative(root *TreeNode) []int {
	resp := []int{}
	if root == nil {
		return resp
	}

	stack := []*TreeNode{}
	curr := root

	for len(stack) > 0 || curr != nil {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		l := len(stack)
		curr = stack[l-1]
		stack = stack[:l-1]
		curr = curr.Right
	}

	return resp
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	left := inorderTraversal(root.Left)
	right := inorderTraversal(root.Right)

	resp := []int{}
	resp = append(resp, left...)
	resp = append(resp, root.Val)
	resp = append(resp, right...)
	return resp
}

func buildTreeFromSlice(values []any) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := &TreeNode{Val: values[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for i < len(values) {
		current := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(values) && values[i] != nil {
			left := &TreeNode{Val: values[i].(int)}
			current.Left = left
			queue = append(queue, left)
		}
		i++

		// Right child
		if i < len(values) && values[i] != nil {
			right := &TreeNode{Val: values[i].(int)}
			current.Right = right
			queue = append(queue, right)
		}
		i++
	}

	return root
}
