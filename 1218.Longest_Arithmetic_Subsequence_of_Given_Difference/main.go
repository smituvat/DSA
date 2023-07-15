package main

import "fmt"

func main() {
	fmt.Print(longestSubsequence([]int{1, 5, 7, 8, 5, 3, 4, 2, 1}, -2))
}

func longestSubsequence(arr []int, diff int) int {
	dp := make(map[int]int)
	answer := 1
	for _, item := range arr {
		dp[item] = dp[item-diff] + 1
		answer = max(answer, dp[item])
	}
	return answer
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
