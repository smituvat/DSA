package main

import "fmt"

func main() {
	s := "ABAB"
	k := 2
	res := characterReplacement(s, k)
	fmt.Print(res)
}

func characterReplacement(s string, k int) int {
	count := make(map[byte]int)
	res := 0

	l := 0
	maxVal := 0
	for r, _ := range s {
		count[s[r]] = count[s[r]] + 1
		maxVal = max(maxVal, count[s[r]])

		if (r-l+1)-maxVal > k {
			count[s[l]] -= 1
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
