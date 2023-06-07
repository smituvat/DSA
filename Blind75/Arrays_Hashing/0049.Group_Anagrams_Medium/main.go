/*
Given an array of strings strs, group the anagrams together. You can return the answer in any order.
An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
eg:
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
*/

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	ans := groupAnagrams(strs)
	fmt.Println(ans)
}

func groupAnagrams(strs []string) [][]string {
	// sort item in array, put them in map which are anagram, append them to a result

	uniqueCheck := make(map[string][]string)

	for _, v := range strs {
		result := sortString(v)
		if _, ok := uniqueCheck[result]; !ok {
			uniqueCheck[result] = []string{}
		}
		uniqueCheck[result] = append(uniqueCheck[result], v)
	}

	res := make([][]string, len(uniqueCheck))

	i := 0

	for _, anagram := range uniqueCheck {
		res[i] = anagram
		i++
	}
	return res
}

func sortString(s string) string {
	strArr := strings.Split(s, "")
	sort.Strings(strArr)
	return strings.Join(strArr, "")
}
