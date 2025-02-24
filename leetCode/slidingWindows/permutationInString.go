package slidingWindows

func checkInclusion(s1 string, s2 string) bool {

	if len(s1) > len(s2) {
		return false
	}

	need := [26]int{}
	for i := range s1 {
		need[s1[i]-'a']++
	}

	have := [26]int{}
	for r := range s2 {
		if r >= len(s1) {
			have[s2[r-len(s1)]-'a']--
		}
		have[s2[r]-'a']++
		if need == have {
			return true
		}
	}
	return false
}
