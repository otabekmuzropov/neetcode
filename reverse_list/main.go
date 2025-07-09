package main

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	node := buildList(nums)
	reverseList(node)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var (
		prev  *ListNode
		cur   = head
		count int
	)

	for cur != nil {
		count++
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}

	return prev
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
