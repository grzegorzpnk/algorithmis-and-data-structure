package ArraysAndHash

import "sort"

//Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

func ContainsDuplicate(nums []int) bool {

	duplicateMap := make(map[int]bool)
	for _, v := range nums {
		if duplicateMap[v] {
			return true
		} else {
			duplicateMap[v] = true
		}
	}
	return false

}

func containsDuplicate2(nums []int) bool {
	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}
