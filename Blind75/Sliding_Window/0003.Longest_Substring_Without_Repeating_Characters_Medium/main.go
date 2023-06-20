package main

import "fmt"

func main() {
	s := "pwwkew"
	res := lengthOfLongestSubstring1(s)
	fmt.Print(res)
}

func lengthOfLongestSubstring(s string) int {
	ss := map[byte]bool{}
	i, ans := 0, 0
	for j := 0; j < len(s); j++ {
		for ss[s[j]] {
			ss[s[i]] = false
			i++
		}
		ss[s[j]] = true
		ans = max(ans, j-i+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// This gives subsequence not substring
func lengthOfLongestSubsequence(s string) int {
	// count := 0
	m := make(map[string]int)
	for _, v := range s {
		s := string(v)
		m[s]++
	}
	return len(m)
}

func lengthOfLongestSubstring1(s string) int {
	charSet := make(map[byte]bool)
	l := 0
	res := 0

	for r, _ := range s {
		for charSet[s[r]] {
			delete(charSet, s[l])
			l++
		}
		charSet[s[r]] = true
		res = max(res, r-l+1)
	}
	return res
}
