package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	five := TreeNode{Val: 5}
	three := TreeNode{Val: 3, Left: &five}
	six := TreeNode{Val: 6}
	fifteen := TreeNode{Val: 15, Left: &three, Right: &six}
	nine := TreeNode{Val: 9}
	eight := TreeNode{Val: 8}
	two := TreeNode{Val: 2, Left: &nine, Right: &eight}
	thirty := TreeNode{Val: 30, Right: &two}

	tn := TreeNode{Val: 10, Right: &thirty, Left: &fifteen}
	target := TreeNode{Val: 9}

	fmt.Print(distanceK(&tn, &target, 2))
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	var (
		res    []int
		find   func(root, target *TreeNode)
		search func(root *TreeNode, distance int)
	)

	mp := make(map[*TreeNode]int)

	find = func(root, target *TreeNode) {
		if root == nil {
			return
		}

		if root == target {
			mp[root] = 0
			return
		}

		find(root.Left, target)
		if val, ok := mp[root.Left]; ok {
			mp[root] = val + 1
			return
		}
		find(root.Right, target)
		if val, ok := mp[root.Right]; ok {
			mp[root] = val + 1
			return
		}
	}

	search = func(root *TreeNode, distance int) {
		if root == nil {
			return
		}

		if val, ok := mp[root]; ok {
			distance = val
		}

		if distance == k {
			res = append(res, root.Val)
		}

		search(root.Left, distance+1)
		search(root.Right, distance+1)
	}
	find(root, target)
	search(root, 0)

	return res
}
