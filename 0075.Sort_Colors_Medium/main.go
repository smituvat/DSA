package main

func main() {
	sortColors([]int{2, 0, 2, 1, 1, 0})
}

func sortColors(nums []int) {
	low, mid := 0, 0
	high := len(nums) - 1

	for mid <= high {
		if nums[mid] == 0 {
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		} else if nums[mid] == 1 {
			mid++
		} else {
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}

}
