package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(repeatedStringMatch("abcd", "cdabcdab"))
}

func repeatedStringMatch(A string, b string) int {
	count := 0
	a := A

	for {
		if strings.Contains(A, b) {
			return count + 1
		} else if len(a)*count > len(b) {
			return -1
		}
		A += a
		count += 1
	}
}

// func repeatedStringMatch(A string, B string) int {
// 	var times = len(B)/len(A) + 2 // times of A can repeat at most.
// 	var ATimes = strings.Repeat(A, times)

// 	if !strings.Contains(ATimes, B) {
// 		return -1
// 	}

// 	for {
// 		ATimes = strings.Repeat(A, times)
// 		if !strings.Contains(ATimes, B) {
// 			return times + 1
// 		}

// 		times--
// 	}
// }
