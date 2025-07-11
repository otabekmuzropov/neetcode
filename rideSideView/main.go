package main

import "fmt"

func main() {
	nums := []any{1, 2, 3, nil, 5, nil, 4}
	node := buildTreeFromSlice(nums)
	fmt.Println(rightSideView(node))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	resp := []int{}

	if root == nil {
		return resp
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		printTree(queue)
		n := len(queue)
		resp = append(resp, queue[n-1].Val)

		for range n {
			curr := queue[0]
			queue = queue[1:]
			if curr == nil {
				continue
			}

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}

			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
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

func printTree(trees []*TreeNode) {
	for _, tree := range trees {
		fmt.Print(tree.Val, " ")
	}
	fmt.Println()
}
