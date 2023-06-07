/*
An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
eg:
Input: s = "anagram", t = "nagaram"
Output: true
Input: s = "rat", t = "car"
Output: false
*/

package main

import "fmt"

func main() {
	s := "smits"
	t := "smita"

	ans := isAnagram(s, t)
	fmt.Print(ans)
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var alpha [26]int

	for i := 0; i < len(s); i++ {
		alpha[s[i]-'a']++
		alpha[t[i]-'a']--
	}

	for j := 0; j < len(alpha); j++ {
		if alpha[j] != 0 {
			return false
		}
	}
	return true

}
