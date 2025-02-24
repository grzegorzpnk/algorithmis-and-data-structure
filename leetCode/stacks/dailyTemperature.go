package stacks

//Input: temperatures = [73,74,75,71,69,72,76,73]
//Output: [1,1,4,2,1,1,0,0]

func DailyTemperatures(temps []int) []int {
	results := make([]int, len(temps))
	stack := make([]int, 0)

	for i, temp := range temps {
		for len(stack) > 0 && temps[stack[len(stack)-1]] < temp {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			results[index] = i - index
		}

		stack = append(stack, i)
	}

	return results
}

//authored solution, slow one
//func dailyTemperatures(temperatures []int) []int {
//
//	results := make([]int, len(temperatures))
//
//	for i := len(temperatures)-2; i >= 0; i-- {
//		results[i]=0
//		for x := i+1; x < len(temperatures); x++{
//			if temperatures[x] > temperatures[i]{
//				results[i]=x-i
//				break
//			}
//
//		}
//
//	}
//	return results
//}
