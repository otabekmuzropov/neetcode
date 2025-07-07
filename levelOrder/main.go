package main

import "fmt"

func main() {
	nums := []any{}
	node := buildTreeFromSlice(nums)
	fmt.Println(levelOrder(node))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	resp := [][]int{}

	if root == nil {
		return resp
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		n := len(queue)
		ith := []int{}
		for range n {
			curr := queue[0]
			queue = queue[1:]
			if curr == nil {
				continue
			}
			ith = append(ith, curr.Val)

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}

			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}

		resp = append(resp, ith)
	}

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
