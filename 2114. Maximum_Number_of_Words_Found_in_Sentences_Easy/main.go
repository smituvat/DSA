// Array and Strings

package main

import (
	"fmt"
)

func main() {

	sentences := []string{"please wait", "continue to fight", "continue to win y"}
	count := mostWordsFound(sentences)
	fmt.Printf("%v\n", count)
	// for _, v := range sentences {
	// 	fmt.Printf("%v\n", v)
	// }
}

func mostWordsFound(sentences []string) int {
	ans := 0
	for _, s := range sentences {
		ans = max(ans, 1+cnsat(s, ' '))
	}
	return ans
}
func cnsat(s string, c rune) int {
	cnt := 0
	for _, ch := range s {
		if ch == c {
			cnt++
		}
	}
	return cnt
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
