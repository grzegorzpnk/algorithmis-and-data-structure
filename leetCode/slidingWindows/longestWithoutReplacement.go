package slidingWindows

func characterReplacement(s string, k int) int {

	left := 0
	checkMap := make(map[byte]int)
	resultCnt := 0
	maxFreqChar := 0

	for right := 0; right < len(s); right++ {

		//ok, mam przesuniety prawy pointer, co z tym zrobic? update checkMapy oraz maxFreqChar
		checkMap[s[right]] += 1
		maxFreqChar = max(maxFreqChar, checkMap[s[right]])

		//jeśli okno jest ważne, to przelicz rezultat, jeśli nie update lewego pointera
		if (right-left+1)-maxFreqChar > k {
			checkMap[s[left]] -= 1
			left += 1
		}
		resultCnt = max(resultCnt, right-left+1)

	}

	return resultCnt

}
