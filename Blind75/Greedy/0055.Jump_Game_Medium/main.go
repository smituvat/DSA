package main

import "fmt"

func main() {
	// nums := []int{2, 3, 1, 1, 4}
	nums1 := []int{3, 2, 1, 0, 4}
	res := canJump(nums1)
	fmt.Print(res)
}

func canJump(nums []int) bool {
	goal := len(nums) - 1

	for i := len(nums) - 2; i >= 0; i-- {
		if i+nums[i] >= goal {
			goal = i
		}
	}
	return goal == 0
}
