package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

type linkedList struct {
	head   *ListNode
	length int
}

func (l *ListNode) prepend(n *ListNode) {
	second := l.Next
	l.Next = n
	l.Next.Next = second
	l.Val++
}

func main() {
	data := ListNode{}
	node1 := &ListNode{Val: 4}
	node2 := &ListNode{Val: 5}
	node3 := &ListNode{Val: 6}

	data.prepend(node1)
	data.prepend(node2)
	data.prepend(node3)

	res := reverseList(&data)
	fmt.Print(&res)
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	reversedListHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return reversedListHead
}

// Iterative version
func reverseLists(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}

	return prev
}
