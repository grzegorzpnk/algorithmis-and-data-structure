package main

func main() {

	myArr := []int{10, 9, 2, 5, 3, 7, 101, 18}
	lengthOfLIS(myArr)

}

func lengthOfLIS(nums []int) int {

	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
	}

	maxVal := 0
	for _, value := range dp {
		if value > maxVal {
			maxVal = value
		}
	}

	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
