package main

import "fmt"

func main() {
	fmt.Print(buddyStrings("ab", "ba"))
}

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	if s == goal {
		return checkTwoBytes(s)
	} else {
		return checkIfBuddy(s, goal)
	}
}

func checkTwoBytes(s string) bool {
	book := make(map[byte]bool, len(s))

	for i := range s {
		if book[s[i]] {
			return true
		}

		book[s[i]] = true
	}

	return false
}

func checkIfBuddy(s string, goal string) bool {
	var count int
	var tmp1, tmp2 byte

	for i := range s {
		if s[i] != goal[i] {
			count++
			if count == 1 {
				tmp1 = s[i]
				tmp2 = goal[i]
			}

			if count == 2 {
				if tmp1 != goal[i] || tmp2 != s[i] {
					return false
				}
			}

			if count > 2 {
				return false
			}
		}
	}

	if count == 1 {
		return false
	} else {
		return true
	}
}
