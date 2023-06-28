package main

import "fmt"

func main() {
	fmt.Print(rob([]int{2, 3, 2}))
}

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	if n == 2 {
		return max(nums[0], nums[1])
	}

	return max(robHelp(nums[:len(nums)-1]), robHelp(nums[1:]))
}

func robHelp(nums []int) int {
	rob1, rob2 := 0, 0

	for _, v := range nums {
		temp := max(v+rob1, rob2)
		rob1 = rob2
		rob2 = temp
	}
	return rob2
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
