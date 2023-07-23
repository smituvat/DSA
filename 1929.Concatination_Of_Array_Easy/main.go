package main

import "fmt"

func main() {
	nums := []int{1, 2, 1}
	fmt.Print(getConcatenation(nums))
}

func getConcatenation(nums []int) []int {
	ans := make([]int, len(nums)*2)
	for i := 0; i < len(nums); i++ {
		ans[i], ans[i+len(nums)] = nums[i], nums[i]
	}

	return ans

}
