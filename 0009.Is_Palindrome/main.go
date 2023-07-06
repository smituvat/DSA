package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print(isPalindrome(-121))
}
func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	i := 0
	j := len(str) - 1
	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}
