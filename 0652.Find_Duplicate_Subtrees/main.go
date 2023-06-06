package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	struTree := &TreeNode{}
	struTree.Val = 2
	struTree.Left.Val = 1
	struTree.Right.Val = 3
	fmt.Print(findDuplicateSubtrees(struTree))
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var ans []*TreeNode
	counter := make(map[string]int)
	var dfs func(root *TreeNode) string
	dfs = func(root *TreeNode) string {
		if root == nil {
			return "#"
		}
		v := strconv.Itoa(root.Val) + "," + dfs(root.Left) + "," + dfs(root.Right)
		counter[v]++
		if counter[v] == 2 {
			ans = append(ans, root)
		}
		return v
	}
	dfs(root)
	return ans
}
