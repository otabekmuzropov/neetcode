package main

import "fmt"

func main() {

	a := &ListNode{1, nil}
	b := &ListNode{2, a}
	c := b
	fmt.Println(&c, &b)
	c = c.Next
	fmt.Println(&c)
	fmt.Println(&b.Next)

	reverseList(b)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	resp := &ListNode{}
	return resp

}
