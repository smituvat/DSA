package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	five := ListNode{Val: 5, Next: &ListNode{Val: 6}}
	four := ListNode{Val: 4, Next: &five}
	three := ListNode{Val: 3, Next: &four}
	six := ListNode{Val: 6, Next: &three}
	two := ListNode{Val: 2, Next: &six}
	head := ListNode{Val: 1, Next: &two}
	fmt.Print(removeElements(&head, 6))
}

func removeElements(head *ListNode, val int) *ListNode {
	// if head == nil {
	// 	return nil
	// } else if head.Val == val {
	// 	return removeElements(head.Next, val)
	// } else {
	// 	head.Next = removeElements(head.Next, val)
	// 	return head
	// }
	dummy := &ListNode{
		Next: head,
	}
	prev := dummy
	cur := head
	for cur != nil {
		if cur.Val == val {
			prev.Next = cur.Next
		} else {
			prev = prev.Next
		}
		cur = cur.Next
	}
	return dummy.Next
}
