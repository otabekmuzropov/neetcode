package main

import (
	"fmt"
)

func main() {
	nums := []any{1, 2, 2, 3, 3, nil, nil, 4, 4}
	root := buildTreeFromSlice(nums)

	fmt.Println(depthoftree(root))
	fmt.Println(isBalanced(root))
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftDepth := depthoftree(root.Left)
	rightDepth := depthoftree(root.Right)
	if leftDepth-rightDepth > 1 || rightDepth-leftDepth > 1 {
		return false
	}
	if !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}

	return true
}

func depthoftree(node *TreeNode) int {
	if node == nil {
		return 0
	}

	maxi := 1 + max(depthoftree(node.Left), depthoftree(node.Right))

	return maxi
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
