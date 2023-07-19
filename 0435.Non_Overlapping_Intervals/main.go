package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	fmt.Print(eraseOverlapIntervals(intervals))
}

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	curInterval := intervals[0]

	res := 1

	for _, val := range intervals[1:] {
		if val[0] >= curInterval[1] {
			curInterval = val
			res++
		}
	}

	return len(intervals) - res
}
