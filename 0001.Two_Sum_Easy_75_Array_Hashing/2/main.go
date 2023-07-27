/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.
eg:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
*/

package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Print(result)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, num := range nums {

		if val, found := m[target-num]; found {
			return []int{val, idx}
		}
		m[num] = idx
	}

	return nil
}
