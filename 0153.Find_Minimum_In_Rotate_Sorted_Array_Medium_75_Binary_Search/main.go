package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 8, 1, 2}
	res := findMin(nums)
	fmt.Print(res)
}

func findMin(nums []int) int {
	res := nums[0]
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] >= nums[0] {
			i = mid + 1
		} else {
			if nums[mid] < res {
				res = nums[mid]
			}
			j = mid - 1
		}
	}
	return res
}
