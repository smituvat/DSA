/*
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
Eg:
Input: nums = [1,2,3,1]
Output: true
*/
package main

import "fmt"

func main() {
	nums := []int{1, 1, 3, 4}
	res := containsDuplicate(nums)
	fmt.Print(res)

}

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for k, v := range nums {
		if _, isunique := m[v]; isunique {
			return true
		}
		m[v] = k
	}
	return false
}
