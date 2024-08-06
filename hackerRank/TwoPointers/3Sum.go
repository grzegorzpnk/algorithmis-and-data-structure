package TwoPointers

import "sort"

//3Sum
//Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
//
//Notice that the solution set must not contain duplicate triplets.
//
//Example 1:
//
//Input: nums = [-1,0,1,2,-1,-4]
//Output: [[-1,-1,2],[-1,0,1]]
//Explanation:
//nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
//nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
//nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
//The distinct triplets are [-1,0,1] and [-1,-1,2].
//Notice that the order of the output and the order of the triplets does not matter.
//Example 2:
//
//Input: nums = [0,1,1]
//Output: []
//Explanation: The only possible triplet does not sum up to 0.
//Example 3:
//
//Input: nums = [0,0,0]
//Output: [[0,0,0]]
//Explanation: The only possible triplet sums up to 0.

func threeSum(nums []int) [][]int {

	var results [][]int
	n := len(nums)
	tripletMap := make(map[[3]int][]int)
	//
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for z := j + 1; z < n; z++ {
				if nums[i]+nums[j]+nums[z] == 0 {
					triplet := [3]int{nums[i], nums[j], nums[z]}
					sort.Ints(triplet[:])
					tripletMap[triplet] = []int{nums[i], nums[j], nums[z]}
				}
			}
		}
	}

	for _, v := range tripletMap {
		results = append(results, v)
	}

	return results

}
