package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	res := search(nums, target)
	fmt.Print(res)
}

func search(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[i] <= nums[mid] {
			if target > nums[mid] || target < nums[i] {
				i = mid + 1
			} else {
				j = mid - 1
			}
		} else {
			if target < nums[mid] || target > nums[j] {
				j = mid - 1
			} else {
				i = mid + 1
			}
		}

	}
	return -1
}
