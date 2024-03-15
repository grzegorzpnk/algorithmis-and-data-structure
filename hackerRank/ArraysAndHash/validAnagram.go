package ArraysAndHash

// https://leetcode.com/problems/valid-anagram/
func isAnagram(s string, t string) bool {

	anagramMap := make(map[rune]int)

	if len(s) != len(t) {
		return false
	}

	for _, rune := range s {
		anagramMap[rune] += 1
	}
	for _, rune := range t {
		if anagramMap[rune] != 0 {
			anagramMap[rune] -= 1
		} else {
			return false
		}
	}
	return true
}
