package main

import "fmt"

func main() {
	word := []string{"ab", "ba", "cc"}
	res := maximumNumberOfStringPairs(word)
	fmt.Print(res)
}

func maximumNumberOfStringPairs(words []string) int {
	count := 0
	s := make(map[string]int)
	for _, v := range words {
		rv := Reverse(v)
		if _, istrue := s[rv]; istrue {
			count++
		}
		s[v]++
	}
	return count
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
