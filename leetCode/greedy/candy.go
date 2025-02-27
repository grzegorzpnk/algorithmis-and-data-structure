package greedy

func candy(ratings []int) int {

	sum := 0
	candies := make([]int, len(ratings))

	//assign everyon 1
	for i := 0; i < len(ratings); i++ {
		candies[i] = 1
	}

	//from left to right
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	//from right to left

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}

	for _, v := range candies {
		sum += v
	}

	return sum
}
