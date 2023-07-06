package main

import "fmt"

func main() {
	fmt.Print(singleNumber([]int{0, 1, 0, 1, 0, 1, 99}))
}

func singleNumber(nums []int) int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	for k, v := range m {
		if v < 3 {
			return k
		}
	}
	return 0
}
