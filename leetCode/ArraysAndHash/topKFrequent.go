package ArraysAndHash

//Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.
//
//Input: nums = [1,1,1,2,2,3], k = 2
//Output: [1,2]
//
//Input: nums = [1], k = 1
//Output: [1]
//ArraysAndHash.TopKFrequent([]int{5, 5, 5, 5, 1, 1, 1, 2, 2, 3, 8, 8, 8, 8, 8, 8, 8}, 2)

func TopKFrequent(nums []int, k int) []int {

	var results []int
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	arr := make([][]int, len(nums)+1)

	for number, quantity := range m {
		arr[quantity] = append(arr[quantity], number)
	}

	for i := len(arr) - 1; i >= 0; i-- {
		for _, v := range arr[i] {
			if k > 0 {
				results = append(results, v)
				k--
			}
		}
	}

	return results
}
