package main

import (
	"fmt"
)

func main() {
	root := buildTreeFromSlice([]any{2, 1, 3, nil, nil, 0, 1})
	fmt.Println(bfs(root))
	fmt.Println(inorder(root))
	fmt.Println(preorder(root))
	fmt.Println(postorder(root))
}

func bfs(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	resp := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		resp = append(resp, cur.Val)

		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}

		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}

	return resp
}

func inorder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	resp := []int{}

	left := inorder(root.Left)
	right := inorder(root.Right)

	resp = append(resp, left...)
	resp = append(resp, root.Val)
	resp = append(resp, right...)

	return resp
}

func preorder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	resp := []int{root.Val}

	left := preorder(root.Left)
	right := preorder(root.Right)

	resp = append(resp, left...)
	resp = append(resp, right...)

	return resp
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
