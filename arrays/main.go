package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, n)

	// Initialize dp with the last row
	copy(dp, triangle[n-1])

	// Bottom-up calculation
	for row := n - 2; row >= 0; row-- {
		for col := 0; col < len(triangle[row]); col++ {
			dp[col] = triangle[row][col] + min(dp[col], dp[col+1])
		}
	}

	return dp[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(minimumTotal(triangle))
}
