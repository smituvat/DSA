package main

import "fmt"

func main() {
	res := shuffle([]int{2, 5, 1, 3, 4, 7}, 3)
	fmt.Print(res)
}

func shuffle(nums []int, n int) []int {
	res := []int{}

	for i := 0; i < n; i++ {
		res = append(res, nums[i], nums[i+n])
	}

	return res
}
