package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	nums := []any{5, 3, 6, 2, 4, nil, 7}
	node := buildTreeFromSlice(nums)
	fmt.Println(bfs(node))
	fmt.Println(bfs((deleteNode(node, 3))))
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			minNode := minValueNode(root.Left)
			root.Val = minNode.Val
			root.Left = deleteNode(root.Left, minNode.Val)
		}
	}

	return root
}

func maxValueNode(root *TreeNode) *TreeNode {
	curr := root
	for curr != nil && curr.Right != nil {
		curr = curr.Right
	}

	return curr
}

func minValueNode(root *TreeNode) *TreeNode {
	curr := root
	for curr != nil && curr.Left != nil {
		curr = curr.Left
	}

	return curr
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
