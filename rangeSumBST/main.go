package main

import (
	"fmt"
)

func main() {
	nums := []any{10, 5, 15, 3, 7, 13, 18, 1, nil, 6}
	node := buildTreeFromSlice(nums)
	// fmt.Println(rangeSumBST(node, 6, 10))
	fmt.Println(rangeSumBSTIterative(node, 6, 10))
	fmt.Println("here")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBSTIterative(root *TreeNode, low, high int) int {
	if root == nil {
		return 0
	}

	sum := 0
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		l := len(stack)
		node := stack[l-1]
		stack = stack[:l-1]

		if node == nil {
			continue
		}

		if node.Val >= low && node.Val <= high {
			sum += node.Val
		}

		if node.Val > low {
			stack = append(stack, node.Left)
		}

		if node.Val < high {
			stack = append(stack, node.Right)
		}
	}

	return sum
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	fmt.Println("value of root: ", root.Val)

	resp := 0

	if root.Val >= low && root.Val <= high {
		resp += root.Val
	}

	left := rangeSumBST(root.Left, low, high)
	right := rangeSumBST(root.Right, low, high)
	resp += left + right

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
