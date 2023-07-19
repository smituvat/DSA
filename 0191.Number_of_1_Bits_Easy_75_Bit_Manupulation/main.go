package main

import "fmt"

func main() {
	fmt.Print(hammingWeight(00000000000000000000000000001011))
}

func hammingWeight(num uint32) int {
	res := 0
	for num > 0 {
		if num&1 == 0 {
			num >>= 1
		} else {
			num &^= 1
			res++
		}
	}
	return res
}
