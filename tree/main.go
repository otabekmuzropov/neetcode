package main

import "fmt"

func main() {
	nums := []any{10, 5, 15, 3, 7, 13, 18, 1, nil, 6}
	node := buildTreeFromSlice(nums)

	fmt.Println(bfsTraversal(node))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bfsTraversal(root *TreeNode) []int {
	resp := make([]int, 0)

	que := []*TreeNode{root}

	for len(que) > 0 {
		node := que[0]
		que = que[1:]

		if node.Left != nil {
			que = append(que, node.Left)
		}

		if node.Right != nil {
			que = append(que, node.Right)
		}

		resp = append(resp, node.Val)

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
