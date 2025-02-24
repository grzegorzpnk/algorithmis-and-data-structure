package ArraysAndHash

import "sort"

//Given an array of strings strs, group the anagrams together. You can return the answer in any order.
//
//An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
//
//
//
//Example 1:
//
//Input: strs = ["eat","tea","tan","ate","nat","bat"]
//Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
//Example 2:
//
//Input: strs = [""]
//Output: [[""]]
//Example 3:
//
//Input: strs = ["a"]
//Output: [["a"]]
//

func groupAnagrams(anagrams []string) [][]string {

	keyToAnagrams := make(map[string][]string)

	for _, anagram := range anagrams {
		chars := []byte(anagram)
		sort.Slice(chars, func(i, j int) bool { return chars[i] < chars[j] })
		key := string(chars)
		keyToAnagrams[key] = append(keyToAnagrams[key], anagram)
	}

	var results [][]string
	for _, v := range keyToAnagrams {
		results = append(results, v)
	}

	return results

}
