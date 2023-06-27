package main

import "fmt"

func main() {
	grid := [][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0}}
	res := numIslands(grid)
	fmt.Print(res)
}

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])

	no_island := 0

	var dfs func(i int, j int)

	dfs = func(i int, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return
		}
		grid[i][j] = 0
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}
	for i, row := range grid {
		for j, point := range row {
			if point == 1 {
				no_island++
				dfs(i, j)
			}
		}
	}
	return no_island
}
