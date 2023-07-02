package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	fmt.Print(merge(intervals))
}

func merge(intervals [][]int) [][]int {
	res := [][]int{}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res = append(res, intervals[0])

	for _, interval := range intervals {
		last := res[len(res)-1]
		if last[1] > interval[0] {
			last[1] = max(last[1], interval[1])
		} else {
			res = append(res, interval)
		}
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
