package main

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}

// brute force
// func rotate(nums []int, k int) {
// 	temp, previous := 0, 0
// 	for i := 0; i < k; i++ {
// 		previous = nums[len(nums)-1]
// 		for j := 0; j < len(nums); j++ {
// 			temp = nums[j]
// 			nums[j] = previous
// 			previous = temp
// 		}
// 	}

// }

func rotate(nums []int, k int) {
	n := len(nums)
	// k %= n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
