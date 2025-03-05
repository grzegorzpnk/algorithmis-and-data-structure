package Math

func plusOne(digits []int) []int {

	//len of digits slice
	n := len(digits)

	//iterate starting from last digit, until we iterato to first (zero-slice) digit
	for i := n - 1; i >= 0; i-- {
		//if the i-th digit is not nine, let's increment it and return digits
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		//otherwise (the currently considered i-th element is 9, so the increment would make it 10, set this digit as a 0)
		digits[i] = 0

	}
	//this is edge case for digits like 999999.. than currently digits looks like: 000000, so we need to append at the beggining 1 and return it
	digits = append([]int{1}, digits...)
	return digits
}
