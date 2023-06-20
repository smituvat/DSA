// Array
// Brute force : 2 iteration then sing max profit O(n^2)

package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	ans := maxProfit(prices)
	fmt.Print(ans)
}

func maxProfit(prices []int) (profit int) {
	buy := prices[0]
	for _, thatdaycost := range prices {
		profit = max(profit, thatdaycost-buy) //0,0,4,4,5,5
		buy = min(buy, thatdaycost)           //7,1,1,1,1,1
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
