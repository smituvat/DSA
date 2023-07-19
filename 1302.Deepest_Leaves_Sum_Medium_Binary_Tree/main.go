package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	eight := TreeNode{Val: 8}
	six := TreeNode{Val: 6, Right: &eight}
	three := TreeNode{Val: 3, Right: &six}
	five := TreeNode{Val: 5}
	seven := TreeNode{Val: 7}
	four := TreeNode{Val: 4, Left: &seven}
	two := TreeNode{Val: 2, Left: &four, Right: &five}
	gRoot := TreeNode{Val: 1, Left: &two, Right: &three}

	fmt.Print(deepestLevelSum(&gRoot))
}

func deepestLevelSum(root *TreeNode) int {
	_, sum := get(root, 0, 0, 0)
	return sum
}

func get(root *TreeNode, level, highLevel, sum int) (int, int) {

	if root.Left != nil {
		highLevel, sum = get(root.Left, level+1, highLevel, sum)
	}

	if root.Right != nil {
		highLevel, sum = get(root.Right, level+1, highLevel, sum)
	}

	if root.Left == nil && root.Right == nil {
		if level > highLevel {
			highLevel = level
			sum = root.Val
		} else if level == highLevel {
			sum += root.Val
		}
	}
	return highLevel, sum
}
