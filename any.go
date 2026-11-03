package main

import "fmt"

func main() {
	tree := buildTreeFromSlice([]any{3, 9, 20, nil, nil, 15, 7})
	fmt.Println("here again")
	fmt.Println(levelOrder(tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	done := [][]int{}
	if root == nil {
		return done
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		l := len(queue)
		ext := make([]int, l)
		for i := range l {
			node := queue[0]
			ext[i] = node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[0:]
		}

		done = append(done, ext)
	}

	return done
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
