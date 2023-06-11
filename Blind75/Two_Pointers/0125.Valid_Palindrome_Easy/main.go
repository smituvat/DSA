/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.
Given a string s, return true if it is a palindrome, or false otherwise.
*/
package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "A man, a plan, a canal: Panama"
	isBool := isPalindrome(s)
	fmt.Print(isBool)
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	arr := []rune(s)

	for i < j {
		left := unicode.ToLower(arr[i])
		right := unicode.ToLower(arr[j])

		if !isLettrrOrDigit(left) {
			i++
			continue
		}

		if !isLettrrOrDigit(right) {
			j--
			continue
		}

		if left != right {
			return false
		}
		i++
		j--
	}
	return true
}

func isLettrrOrDigit(s rune) bool {
	return unicode.IsLetter(s) || unicode.IsDigit(s)
}
