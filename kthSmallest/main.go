package main

import (
	"fmt"
)

func main() {
	nums := []any{5, 3, 6, 2, 4, nil, nil, 1}
	n := 3
	node := buildTreeFromSlice(nums)
	fmt.Println(kthSmallest(node, n))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	values := inorderN(root, k)
	fmt.Println(values)
	return values[k-1]
}

func inorderN(root *TreeNode, n int) []int {
	var result []int
	var count int
	inorderLimited(root, n, &count, &result)
	return result
}

func inorderLimited(node *TreeNode, n int, count *int, result *[]int) {
	if node == nil || *count >= n {
		return
	}
	fmt.Println("here again", node.Val)

	inorderLimited(node.Left, n, count, result)

	if *count < n {
		*result = append(*result, node.Val)
		*count++
	}

	inorderLimited(node.Right, n, count, result)
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
