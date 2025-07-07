package main

import (
	"fmt"
)

func main() {
	nums := []any{1, 2, 3}
	node := buildTreeFromSlice(nums)
	fmt.Println(countLeavesIterative(node))
	fmt.Println(countLeavesRecursion(node))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countLeavesIterative(root *TreeNode) int {
	if root == nil {
		return 0
	}

	count := 0
	que := []*TreeNode{root}
	for len(que) > 0 {
		node := que[0]
		que = que[1:]
		if node == nil {
			continue
		}

		if node.Left == nil && node.Right == nil {
			count++
		}

		que = append(que, node.Left)
		que = append(que, node.Right)
	}

	return count
}

func countLeavesRecursion(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	return countLeavesRecursion(root.Left) + countLeavesRecursion(root.Right)
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
