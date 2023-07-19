package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxDiameter int

func main() {
	// five := TreeNode{Val: 5}
	// four := TreeNode{Val: 4}
	// three := TreeNode{Val: 3}
	// two := TreeNode{Val: 2, Left: &four, Right: &five}
	// gNode := TreeNode{Val: 1, Left: &two, Right: &three}
	two := TreeNode{Val: 2}
	gNode := TreeNode{Val: 1, Left: &two}
	fmt.Print(diameterOfBinaryTree(&gNode))
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter = 0
	ans := maxDepth(root)
	fmt.Println(ans)
	return maxDiameter
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	// this is the diameter! as the edges between root and left sub tree and root & right sub tree
	maxDiameter = max(maxDiameter, left+right)

	// this is the max depth
	return 1 + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
