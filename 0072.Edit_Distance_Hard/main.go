// Dynamic programming example

/*
	nil r o s

nil 0 1 2 3
h  1 1 2 3
o  2 2 1 2
r  3 2 2 2
s  4 3 3 2
e  5 4 4 *3* ans
*/
package main

import "fmt"

func main() {
	word1 := "horse"
	word2 := "ros"
	ans := minDistance(word1, word2)
	fmt.Print(ans)

}

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
