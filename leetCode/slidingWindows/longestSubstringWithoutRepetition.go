package slidingWindows

func lengthOfLongestSubstring(s string) int {

	max_length := 0
	l := 0
	checkMap := make(map[byte]bool)
	for r := 0; r < len(s); r++ {
		for checkMap[s[r]] == true {
			checkMap[s[l]] = false
			l++
		}

		len := r - l + 1
		max_length = max(len, max_length)
		checkMap[s[r]] = true

	}
	return max_length
}

//func lengthOfLongestSubstring2(s string) int {
//
//	max_length := 0
//	l := 0
//	checkMap := make(map[byte]bool)
//
//	for r := 0; r < len(s); r++ {
//
//			checkMap[s[l]] = false
//			l++
//		}
//
//		len := r - l + 1
//		max_length = max(len, max_length)
//		checkMap[s[r]] = true
//
//	}
//	return max_length
//}
