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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
