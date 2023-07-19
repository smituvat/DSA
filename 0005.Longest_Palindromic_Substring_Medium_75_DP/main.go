package main

import "fmt"

func main() {
	fmt.Print(longestPalindrome("babad"))
}

func longestPalindrome(s string) string {
	res := ""
	resLen := 0

	for i := range s {
		// odd palindrome result
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if (r - l + 1) > resLen {
				res = s[l : r+1]
				resLen = r - l + 1
			}
			l--
			r++
		}

		// even
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if (r - l + 1) > resLen {
				res = s[l : r+1]
				resLen = r - l + 1
			}
			l--
			r++
		}
	}
	return res
}
