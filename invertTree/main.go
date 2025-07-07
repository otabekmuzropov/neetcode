package main

import "fmt"

func main() {
	nums := []any{4, 2, 7, 1, 3, 6, 9}
	node := buildTreeFromSlice(nums)
	node1 := buildTreeFromSlice(nums)
	fmt.Println(bfs(invertTree(node)))
	fmt.Println(bfs(invertTreeRecursive(node1)))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	que := []*TreeNode{root}
	for len(que) > 0 {
		node := que[0]
		que = que[1:]
		if node == nil {
			continue
		}

		node.Left, node.Right = node.Right, node.Left
		que = append(que, node.Left)
		que = append(que, node.Right)
	}

	return root
}

func invertTreeRecursive(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTreeRecursive(root.Left)
	invertTreeRecursive(root.Right)

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
