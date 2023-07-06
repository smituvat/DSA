package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{8, 8, 7, 7, 7}))
}

// func majorityElement(nums []int) int {
// 	mp := make(map[int]int)
// 	max, val := 0, 0
// 	for _, v := range nums {
// 		mp[v]++
// 	}
// 	for k, v := range mp {
// 		if v > max {
// 			max = v
// 			val = k
// 		}
// 	}
// 	return val
// }

func majorityElement(nums []int) int {
	candidate := nums[0]
	count := 1
	for _, num := range nums[1:] {
		if num == candidate {
			count++
		} else {
			count--
		}
		if count == 0 {
			candidate = num
			count = 1
		}
	}
	return candidate
}
