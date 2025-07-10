package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	nums := []any{4, 2, 7, 1, 3}
	node := buildTreeFromSlice(nums)
	fmt.Println(bfs(node))
	fmt.Println(bfs((insertIntoBST(node, 5))))
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	fmt.Println(bfs(root))
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}

	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	} else {
		root.Left = insertIntoBST(root.Left, val)
	}

	return root
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
