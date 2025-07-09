package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	node := buildList(nums)
	fmt.Println(pairSum(node))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	maxi := 0

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev *ListNode
	curr := slow
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}

	for prev != nil {
		maxi = max(maxi, prev.Val+head.Val)
		prev = prev.Next
		head = head.Next
	}

	return maxi
}

// Helper to create list from slice
func buildList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	curr := head
	for _, v := range values[1:] {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return head
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val, " ")
		node = node.Next
	}
	fmt.Println()
}
