package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	data := ListNode{
		Val: 5,
	}
	data.Val = 4
	data.Next.Val = 5
	res := reverseList(&data)
	fmt.Print(res)
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
