package main

import "fmt"

func main() {
	fmt.Print(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}

func findKthLargest(nums []int, k int) int {
	return recursion(nums, k, 0, len(nums)-1)
}

func recursion(nums []int, k int, low, high int) int {
	pivot := partition(nums, low, high)

	if pivot == high-k+1 {
		return nums[pivot]
	}

	if pivot > high-k+1 {
		return recursion(nums, k-(high-pivot+1), low, pivot-1)
	}

	return recursion(nums, k, pivot+1, high)
}

func partition(nums []int, low, high int) int {
	pivot, left := high, low

	for i := low; i <= high; i++ {
		if nums[i] < nums[pivot] {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}

	nums[left], nums[pivot] = nums[pivot], nums[left]

	return left
}
