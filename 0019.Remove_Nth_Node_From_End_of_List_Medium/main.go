package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	four := ListNode{Val: 4}
	three := ListNode{Val: 3, Next: &four}
	six := ListNode{Val: 6, Next: &three}
	two := ListNode{Val: 2, Next: &six}
	head := ListNode{Val: 1, Next: &two}
	removeNthFromEnd(&head, 2)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	first_node := head
	second_node := head
	i := 0
	for first_node != nil {
		if i > n {
			second_node = second_node.Next
		}
		first_node = first_node.Next
		i += 1
	}

	if n == i {
		head = head.Next
	} else {
		second_node.Next = second_node.Next.Next
	}
	return head
}
