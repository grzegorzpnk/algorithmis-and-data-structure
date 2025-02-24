package TwoPointers

func isPalindrome(s string) bool {

	i := 0
	j := len(s) - 1

	for i < j {
		if isPunctual(s[i]) {
			i++
			continue
		}

		if isPunctual(s[j]) {
			j--
			continue
		}

		if toSmallerCase(s[i]) == toSmallerCase(s[j]) {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func toSmallerCase(c byte) byte {

	if c >= 65 && c <= 90 {
		return c + 32
	}
	return c

}

func isPunctual(c byte) bool {
	if (c >= 32 && c <= 47) || (c >= 58 && c <= 64) || (c >= 91 && c <= 96) || (c >= 123 && c <= 126) {
		return true
	}
	return false
}
