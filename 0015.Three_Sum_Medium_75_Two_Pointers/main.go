/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
Notice that the solution set must not contain duplicate triplets.
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	fmt.Print(res)
}

func threeSum(nums []int) [][]int {
	// first sort the array as we shouldnt have duplicate values
	sort.Ints(nums)

	n := len(nums)
	var result [][]int

	for a := 0; a < n-2; a++ {

		// condition to eleminate duplicate vals
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		b := a + 1
		c := n - 1
		for b < c {
			sum := nums[a] + nums[b] + nums[c]
			if sum == 0 {
				result = append(result, []int{nums[a], nums[b], nums[c]})

				c--
				for b < c && nums[c] == nums[c+1] {
					c--
				}
			} else if sum > 0 {
				c--
			} else {
				b++
			}
		}
	}
	return result
}
