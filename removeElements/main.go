package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := buildList([]int{7, 7, 7, 7})
	middle := removeElements(head, 7)
	printList(middle)
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	cur := head

	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	if head.Val == val {
		head = head.Next
	}

	return head
}

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
