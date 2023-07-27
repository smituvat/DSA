package main

func main() {
	nextPermutation([]int{2, 3, 1})
}

func nextPermutation(nums []int) {
	index := -1

	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			index = i
			break
		}
	}
	if index == -1 {
		reverse(nums)
		return
	}

	for i := len(nums) - 2; i > 0; i-- {
		if nums[index] < nums[i] {
			nums[index], nums[i] = nums[i], nums[index]
			break
		}
	}

	reverse(nums[index+1:])

}

func reverse(numbers []int) []int {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}
