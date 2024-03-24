package TwoPointers

func TwoSum(numbers []int, target int) []int {

	for i := 0; i < len(numbers)-2; i++ {
		for j := 1; j < len(numbers)-1; j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}
	return []int{}
}

func TwoSum2(numbers []int, target int) []int {

	i := 0
	j := len(numbers) - 1

	for i < j {

		if numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		}

		if numbers[i]+numbers[j] > target {
			j--
		}

		if numbers[i]+numbers[j] < target {
			i++
		}

	}
	return []int{}

}
