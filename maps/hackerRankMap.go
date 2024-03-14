package maps

import "sort"

// LeetCode issue number 18 -> Four sum
//
//Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:
//
//0 <= a, b, c, d < n
//a, b, c, and d are distinct.
//nums[a] + nums[b] + nums[c] + nums[d] == target
//You may return the answer in any order.
//
//
//
//Example 1:
//
//Input: nums = [1,0,-1,0,-2,2], target = 0
//Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
//Example 2:
//
//Input: nums = [2,2,2,2,2], target = 8
//Output: [[2,2,2,2]]
//
//
//Constraints:
//
//1 <= nums.length <= 200
//-109 <= nums[i] <= 109
//-109 <= target <= 109

func fourSum(nums []int, target int) [][]int {
	result := [][]int{}
	n := len(nums)

	if n < 4 {
		return result
	}

	sort.Ints(nums)

	numsMap := make(map[int][][2]int)

	// Pre-compute the sum of pairs and store them in a map
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			sum := nums[i] + nums[j]
			numsMap[sum] = append(numsMap[sum], [2]int{i, j})
		}
	}

	// Iterate over pairs (nums[i], nums[j]) and find pairs (nums[k], nums[l]) such that nums[i] + nums[j] + nums[k] + nums[l] == target
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			remaining := target - (nums[i] + nums[j])

			// Check if there are pairs in numsMap that can contribute to the remaining sum
			if pairs, found := numsMap[remaining]; found {
				for _, pair := range pairs {
					if pair[0] > j {
						result = append(result, []int{nums[i], nums[j], nums[pair[0]], nums[pair[1]]})
					}
				}
			}
		}
	}

	return result
}

func threeSums(nums []int) [][]int {
	var result [][]int
	tripleMaps := make(map[[3]int][]int)

	for i := 0; i < len(nums)-2; i++ {
		for x := i + 1; x < len(nums)-1; x++ {
			for z := x + 1; z < len(nums); z++ {
				if nums[i]+nums[x]+nums[z] == 0 {
					triple := [3]int{nums[i], nums[x], nums[z]}
					sort.Ints(triple[:])
					tripleMaps[triple] = []int{nums[i], nums[x], nums[z]}
				}
			}
		}
	}

	for _, v := range tripleMaps {
		result = append(result, v)
	}
	return result
}
