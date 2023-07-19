package main

import "fmt"

func main() {
	fmt.Print(shortestPalindrome("aacecaaa"))
}

func shortestPalindrome(s string) string {
	i, j := 0, len(s)-1
	k := 0
	for i < j {
		if s[i] == s[j] {
			i++
			j--
			continue
		} else {
			s = s[:k] + string(s[j]) + s[i:]
			i++
			k++
		}
	}
	return s
}
