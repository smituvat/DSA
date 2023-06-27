package main

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

func main() {
	base := initBTree()
	fmt.Println("pre order")
	base.preOrder()
}

func initBTree() *Node {
	five := Node{val: 5}
	three := Node{val: 3, left: &five}
	six := Node{val: 6}
	fifteen := Node{val: 15, left: &three, right: &six}
	nine := Node{val: 9}
	eight := Node{val: 8}
	two := Node{val: 2, left: &nine, right: &eight}
	thirty := Node{val: 30, right: &two}
	return &Node{val: 10, right: &thirty, left: &fifteen}
}

func (n *Node) preOrder() {
	if n != nil {
		fmt.Println(n.val)
		n.left.preOrder()
		n.right.preOrder()
	}
}
