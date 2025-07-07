package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	node := &ListNode{}
	temp := node
	for list1 != nil && list2 != nil {
		newNode := &ListNode{}
		if list1.Val < list2.Val {
			newNode.Val = list1.Val
			list1 = list1.Next
		} else {
			newNode.Val = list2.Val
			list2 = list2.Next
		}

		temp.Next = newNode
		temp = newNode
	}

	for list1 != nil {
		newNode := &ListNode{Val: list1.Val}
		temp.Next = newNode
		temp = newNode
		list1 = list1.Next
	}
	for list2 != nil {
		newNode := &ListNode{Val: list2.Val}
		temp.Next = newNode
		temp = newNode
		list2 = list2.Next
	}

	return node.Next
}

func main() {
	// Example list: [1, 2, 3, 4, 5]
	node1 := buildList([]int{1, 2, 4})
	node2 := buildList([]int{1, 3, 4})
	middle := mergeTwoLists(node1, node2)
	fmt.Println(node1 == node2)
	printList(middle)
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val, " ")
		node = node.Next
	}
	fmt.Println()
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
