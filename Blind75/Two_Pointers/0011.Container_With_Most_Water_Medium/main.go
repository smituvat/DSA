package main

import "fmt"

func main() {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	result := maxArea(input)
	fmt.Print(result)
}

// Brute Force
// func maxArea(height []int) int {
// 	maxVal := 0
// 	for i := 0; i < len(height); i++ {
// 		for j := i + 1; j < len(height); j++ {
// 			area := (j - i) * min(height[i], height[j])
// 			if area > maxVal {
// 				maxVal = area
// 			}
// 		}
// 	}
// 	return maxVal
// }

func maxArea(height []int) int {
	maxVal := 0
	i := 0
	j := len(height) - 1

	for i < j {
		area := (j - i) * min(height[i], height[j])
		if area > maxVal {
			maxVal = area
		}
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return maxVal
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
