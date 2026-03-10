package main

import "fmt"

func main() {
	arr := []int{-10, -3, 0, 5, 9}
	root := sortedArrayToBST(arr)
	printByLevel(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	low, high := 0, len(nums)-1
	for low <= high {

	}

	return nil
}

func printByLevel(root *TreeNode) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)

		// Print one full level
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			fmt.Printf("%d ", node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		fmt.Println() // move to next line for next level
	}
}
