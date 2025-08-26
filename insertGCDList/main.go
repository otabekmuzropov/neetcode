package main

import "fmt"

func main() {
	// Example list: [1, 2, 3, 4, 5]
	head := buildList([]int{7})
	middle := insertGreatestCommonDivisors(head)
	printNode(middle)
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head

	for cur.Next != nil {
		temp := cur.Next
		newNode := &ListNode{
			Val:  gcd(cur.Val, cur.Next.Val),
			Next: temp,
		}

		cur.Next = newNode
		cur = cur.Next.Next
	}

	return head
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
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

func printNode(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " -> ")
		head = head.Next
	}
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}

	return gcd(b%a, a)
}
