package ArraysAndHash

//
//func LongestConsecutive(nums []int) int {
//
//	//numsSet := make(map[int]bool)
//	//for _, n := range nums {
//	//	numsSet[n] = true
//	//}
//	//
//	//maxSequenceLen := 0
//	//
//	//for n := range numsSet {
//	//
//	//	_, ok := numsSet[n-1]
//	//	if !ok {
//	//		SequenceLen := 1
//	//		for i:=n; numsSet[n+SequenceLen]; i++ {
//	//				SequenceLen++
//	//			}
//	//		}
//	//	}
//	//
//	//}
//
//	//return maxSequenceLen
//}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

func LongestConsecutive(nums []int) int {

	var result = 0
	consecMap := make(map[int]bool)

	for _, v := range nums {
		consecMap[v] = true
	}

	cnt := 1
	for _, v := range nums {
		if consecMap[v-1] {
			continue
		} else {
			for i := 1; i < len(nums); i++ {
				if consecMap[v+i] {
					cnt += 1
				} else {
					result = max(result, cnt)
					cnt = 0
					break
				}
			}
		}
	}

	result = max(result, cnt)

	return result

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
