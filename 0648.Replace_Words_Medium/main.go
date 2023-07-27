package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	dictionary := []string{"cat", "bat", "rat"}
	fmt.Print(replaceWords(dictionary, "the cattle was rattled by the battery"))
}

func replaceWords(dictionary []string, sentence string) string {
	tmp_s := strings.Split(sentence, " ")
	sort.Strings(dictionary)
	for i, j := range tmp_s {
		for _, s := range dictionary {
			if len(j) >= len(s) && s == j[0:len(s)] {
				tmp_s[i] = s
				break
			}
		}
	}
	return strings.Join(tmp_s, " ")
}
