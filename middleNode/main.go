package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Find middle node using slow and fast pointers
func findMiddle(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
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

func main() {
	// Example list: [1, 2, 3, 4, 5]
	head := buildList([]int{1, 2, 3, 4, 5, 6})
	middle := findMiddle(head)
	fmt.Println("Middle Node Value:", middle.Val)
}
