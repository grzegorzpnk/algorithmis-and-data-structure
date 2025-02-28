package dynamicProgramming

func houseRobber(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) < 3 {
		return max(nums[0], nums[1])
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[1], nums[0])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(nums[i]+dp[i-2], dp[i-1])
	}

	return max(dp[len(nums)-1])

}
