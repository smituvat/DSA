package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	res := maxSubArray(nums)
	fmt.Print(res)
}

func maxSubArray(nums []int) int {
	result := nums[0]
	total := 0

	for i := 0; i < len(nums); i++ {
		total += nums[i]
		result = max(result, total)

		if total < 0 {
			total = 0
		}
	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
