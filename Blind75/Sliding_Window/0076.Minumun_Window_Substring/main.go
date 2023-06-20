package main

import "fmt"

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	res := minWindow(s, t)
	fmt.Print(res)
}

func minWindow(s string, t string) string {
	start, end := 0, 0
	targetF := make(map[uint8]int)
	curretF := make(map[uint8]int)
	disCharCount := 0
	minSubStr := ""

	for v := range t {
		targetF[t[v]]++
	}

	for end < len(s) {
		curretF[s[end]]++

		if targetF[s[end]] != 0 && targetF[s[end]] == curretF[s[end]] {
			disCharCount++
		}

		for disCharCount == len(targetF) {
			if minSubStr == "" {
				minSubStr = s[start : end+1]
			}
			if end-start+1 < len(minSubStr) {
				minSubStr = s[start : end+1]
			}
			curretF[s[start]]--
			if curretF[s[start]] < targetF[s[start]] {
				disCharCount--
			}
			start++
		}
		end++
	}
	return minSubStr
}
