package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
	Prev *ListNode
}

func main() {
	head1 := buildList([]int{1, 2, 3, 4, 5, 6})
	head2 := buildList([]int{1, 2, 3, 4, 5, 6})
	middle := getIntersectionNode(head1, head2)
	fmt.Println("Middle Node Value:", middle.Val)
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var (
		lastA *ListNode
		lastB *ListNode
	)

	lastA = lastB

	for headA.Next != nil {
		lastA = headA
		headA = headA.Next
	}

	for headB.Next != nil {
		lastB = headB
		headB = headB.Next
	}

	return lastA
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
