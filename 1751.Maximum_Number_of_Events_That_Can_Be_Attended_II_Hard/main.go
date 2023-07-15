package main

import (
	"fmt"
	"sort"
)

func main() {
	events := [][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 10}}
	fmt.Print(maxValue(events, 2))
}

const (
	START = 0
	END   = 1
	VALUE = 2
)

func maxValue(events [][]int, k int) int {
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, len(events)+1)
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i][START] < events[j][START]
	})
	searchCache := map[int]int{}
	for curIdx := len(events) - 1; curIdx >= 0; curIdx-- {
		e := events[curIdx]

		nextIdx, ok := searchCache[e[END]]
		if !ok {
			nextIdx = sort.Search(len(events), func(i int) bool {
				return e[END] < events[i][START]
			})
			searchCache[e[END]] = nextIdx
		}

		for kCnt := 1; kCnt <= k; kCnt++ {
			dp[kCnt][curIdx] = max(
				dp[kCnt][curIdx+1],
				events[curIdx][VALUE]+dp[kCnt-1][nextIdx],
			)
		}
	}
	return dp[k][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
