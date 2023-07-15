package main

import "fmt"

func main() {
	fmt.Print(longestConsecutive([]int{400, 100, 4, 3, 22, 23, 2, 2, 1}))
}

func longestConsecutive(nums []int) int {
	storeNum := map[int]bool{}

	for _, num := range nums {
		storeNum[num] = true
	}

	res := 0

	for _, num := range nums {
		if storeNum[num-1] {
			continue
		}

		seq := 1
		tp := num + 1

		for storeNum[tp] {
			seq++
			tp++
		}

		if seq > res {
			res = seq
		}
	}
	return res
}
