package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(reverseWords("a good   example"))
}

func reverseWords(s string) string {
	ans := strings.Fields(s)
	// fmt.Print(ans)
	// output := strings.Split(strings.Trim(s, " "), " ")
	i := 0
	j := len(ans) - 1
	for i < j {
		temp := ans[i]
		ans[i] = ans[j]
		ans[j] = temp

		i++
		j--
	}
	return strings.Join(ans, " ")
}
