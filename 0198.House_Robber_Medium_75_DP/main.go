package main

import "fmt"

func main() {
	fmt.Print(rob([]int{2, 7, 9, 3, 1}))
}

func rob(nums []int) int {
	rob1, rob2 := 0, 0
	// rob1, rob2, n , n+ 1, .....
	for _, v := range nums {
		temp := max(v+rob1, rob2)
		rob1 = rob2
		rob2 = temp
	}
	return rob2
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
