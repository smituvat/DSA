// Practice again
package main

func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	curr := make([]int, 0)
	var dfs func(i int, curSum int, cur []int)

	dfs = func(i int, curSum int, cur []int) {
		if curSum == target {
			ans = append(ans, append([]int{}, curr...))
			return
		}
		if curSum > target {
			return
		}
		for j := i; j < len(candidates); j++ {
			curr = append(curr, candidates[j])
			dfs(j, curSum+candidates[j], curr)
			curr = curr[:len(curr)-1]
		}
	}
	dfs(0, 0, curr)
	return ans
}
