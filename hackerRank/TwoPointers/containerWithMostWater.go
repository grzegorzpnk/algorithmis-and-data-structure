package TwoPointers

func maxArea(height []int) int {

	maxResult := 0

	for i := 0; i < len(height)-1; i++ {
		for j := len(height) - 1; j > i; j-- {
			currentSize := (j - i) * min(height[i], height[j])
			maxResult = max(maxResult, currentSize)
		}
	}

	return maxResult
}

func maxArea2(height []int) int {
	left := 0
	right := len(height) - 1
	mArea := 0
	for left < right {
		//var minHeight int
		//if height[left] < height[right] {
		//	minHeight = height[left]
		//}else {
		//	minHeight = height[right]
		//}
		minHeight := min(height[left], height[right])
		mArea = max(mArea, minHeight*(right-left))
		//if mArea < (minHeight * (right - left)) {
		//	mArea = minHeight * (right - left)
		//}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return mArea
}
