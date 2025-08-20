package main

import (
	"fmt"
)

func main() {
	nums1 := []any{}
	nums2 := []any{}
	p := buildTreeFromSlice(nums1)
	q := buildTreeFromSlice(nums2)
	fmt.Println(p == nil, q == nil)

	fmt.Println(isSameTree(p, q))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if !isEqual(p, q) {
		return false
	}

	pque := []*TreeNode{p}
	qque := []*TreeNode{q}

	for len(qque) > 0 && len(pque) > 0 {
		lp := len(pque)
		lq := len(qque)
		if lp != lq {
			return false
		}

		for range lq {
			pnode := pque[0]
			pque = pque[1:]
			qnode := qque[0]
			qque = qque[1:]

			if !isEqual(pnode.Left, qnode.Left) {
				return false
			}

			if !isEqual(pnode.Right, qnode.Right) {
				return false
			}

			if pnode.Left != nil {
				pque = append(pque, pnode.Left)
				qque = append(qque, qnode.Left)
			}

			if qnode.Right != nil {
				pque = append(pque, pnode.Right)
				qque = append(qque, qnode.Right)
			}

		}
	}

	return true
}

func isEqual(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil && q != nil {
		return false
	} else if q == nil && p != nil {
		return false
	}

	return p.Val == q.Val
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
