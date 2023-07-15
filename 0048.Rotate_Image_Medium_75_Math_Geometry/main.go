package main

import "fmt"

func main() {
	matrix := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	fmt.Println(matrix)
	rotate(matrix)
}

func rotate(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])

	for i := 0; i < n; i++ {
		// transpose
		for j := i; j < m; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
		// flip row
		for j, k := 0, m-1; j < k; j, k = j+1, k-1 {
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
		}
	}
	fmt.Println(matrix)
}
