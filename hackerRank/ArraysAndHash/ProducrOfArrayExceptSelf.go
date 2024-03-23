package ArraysAndHash

//func productExceptSelf(nums []int) []int {
//
//	results := make([]int, 0)
//	for i := 0; i < len(nums); i++ {
//		product := 1
//		for j := 0; j < len(nums); j++ {
//			if i != j {
//				product = product * nums[j]
//			}
//		}
//		results = append(results, product)
//
//	}
//	return results
//}

func ProductExceptSelf(nums []int) []int {
	rightProduct := make([]int, len(nums))
	results := make([]int, len(nums))

	results[0] = 1
	rightProduct[len(nums)-1] = 1

	//calculate product of left side
	for i := 1; i < len(nums); i++ {
		results[i] = results[i-1] * nums[i-1]
	}

	//calculate product of ritght side
	for i := len(nums) - 2; i >= 0; i-- {
		rightProduct[i] = nums[i+1] * rightProduct[i+1]
		results[i] = rightProduct[i] * results[i]
	}
	return results
}
