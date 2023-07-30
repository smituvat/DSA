package main

// Strings

import (
	"fmt"
	"strings"
)

func main() {
	arr1 := []string{"smita", "supi"}
	arr2 := []string{"supi", "smita"}
	ans := arrayStringsAreEqual(arr1, arr2)
	if ans {
		fmt.Printf("equal")
	} else {
		fmt.Printf("not equal")
	}
}

func arrayStringsAreEqual(arr1 []string, arr2 []string) bool {
	return strings.Join(arr1, "") == strings.Join(arr2, "")
}
