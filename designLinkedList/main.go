package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
	Prev *ListNode
}

type MyLinkedList struct {
	Head *ListNode
	Tail *ListNode
	L    int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		Head: nil,
		Tail: nil,
		L:    0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.L {
		return -1
	}

	curr := this.Head
	for i := 1; i <= index; i++ {
		curr = curr.Next
	}

	return curr.Val
}
func (this *MyLinkedList) AddAtHead(val int) {
	prev := &ListNode{
		Val:  val,
		Prev: nil,
		Next: this.Head,
	}

	if this.Head == nil {
		this.Tail = prev
	} else {
		this.Head.Prev = prev
	}

	this.Head = prev
	this.L++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &ListNode{
		Val:  val,
		Prev: this.Tail,
	}
	if this.Tail != nil {
		this.Tail.Next = node
	} else {
		this.Head = node
	}

	this.Tail = node
	this.L++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.L {
		return
	}

	if index == this.L {
		this.AddAtTail(val)
		return
	}

	if index == 0 {
		this.AddAtHead(val)
		return
	}

	curr := this.Head
	for range index {
		curr = curr.Next
	}

	prev := curr.Prev
	newNode := &ListNode{
		Val:  val,
		Next: curr,
		Prev: prev,
	}
	prev.Next = newNode
	curr.Prev = newNode
	this.L++

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.L {
		return
	}

	curr := this.Head
	for range index {
		curr = curr.Next
	}

	prev := curr.Prev
	next := curr.Next
	if prev != nil {
		prev.Next = next
	} else {
		this.Head = next
	}

	if next != nil {
		next.Prev = prev
	} else {
		this.Tail = prev
	}

	this.L--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
