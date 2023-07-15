/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
An input string is valid if:
Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
Eg:
Input: s = "(]"
Output: false
*/
package main

import "fmt"

func main() {
	s := "()"
	res := isValid(s)
	fmt.Print(res)
}

func isValid(s string) bool {
	// have key val pair in maps

	pairs := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := make([]byte, 0)

	for _, char := range []byte(s) {
		pair, ok := pairs[char]

		if !ok {
			stack = append(stack, char)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		if stack[len(stack)-1] != pair {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0

}
