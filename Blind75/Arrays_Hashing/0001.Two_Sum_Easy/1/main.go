// each input would have exactly one solution, and you may not use the same element twice.
// Hash map

package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Print(result)
}

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i := 0; ; i++ {
		x := nums[i]
		y := target - x
		if j, ok := m[y]; ok {
			return []int{j, i}
		}
		m[x] = i
	}
}
