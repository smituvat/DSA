package main

import "fmt"

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(numbers, target)
	fmt.Print(res)
}

func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1
	fmt.Println(len(numbers))
	for i < j {
		n := numbers[i] + numbers[j]
		if n == target {
			return []int{i + 1, j + 1}
		}
		if n > target {
			j--
		} else {
			i++
		}
	}
	return []int{0, 0}
}
