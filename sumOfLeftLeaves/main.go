package main

import (
	"fmt"
)

func main() {
	nums := []any{1, 2, 3, 4, 5}
	node := buildTreeFromSlice(nums)
	fmt.Println(sumOfLeftLeaves(node))
	fmt.Println("here")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	que := []*TreeNode{root}
	sum := 0

	for len(que) > 0 {
		curr := que[0]
		que = que[1:]

		if curr == nil {
			continue
		}

		if curr.Left != nil {
			que = append(que, curr.Left)
			if curr.Left.Left == nil && curr.Left.Right == nil {
				sum += curr.Left.Val
			}
		}

		if curr.Right != nil {
			que = append(que, curr.Right)
		}
	}

	return sum
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
