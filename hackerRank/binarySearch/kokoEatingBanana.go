package binarySearch

func MinEatingSpeed(piles []int, h int) int {

	for i := 1; i < 1000; i++ {

		sum := 0
		for _, v := range piles {

			if v >= i {
				sum += v / i

				if v%i != 0 {
					sum += 1
				}
			} else {
				sum += i / v
			}

		}

		if sum < h {
			return i
		}

	}
	return 1

}
