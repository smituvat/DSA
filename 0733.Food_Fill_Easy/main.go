package main

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {

	h := len(image)
	w := len(image[0])

	var dfs func(int, int, int)
	dfs = func(r, c, srcColor int) {
		if r < 0 || r >= h || c < 0 || c >= w || image[r][c] == newColor || image[r][c] != srcColor {
			return
		}
		image[r][c] = newColor
		dfs(r-1, c, srcColor)
		dfs(r+1, c, srcColor)
		dfs(r, c-1, srcColor)
		dfs(r, c+1, srcColor)

		return
	}
	dfs(sr, sc, image[sr][sc])
	return image
}
