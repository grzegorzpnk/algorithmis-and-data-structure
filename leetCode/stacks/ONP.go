package stacks

import "strconv"

func EvalRPN(tokens []string) int {
	ONPStack := make([]int, 0)
	var result int

	for _, token := range tokens {
		if token != "-" && token != "+" && token != "*" && token != "/" {
			tokenInt, _ := strconv.Atoi(token)
			ONPStack = append(ONPStack, tokenInt)
		} else {
			switch token {
			case "-":
				result = ONPStack[len(ONPStack)-2] - ONPStack[len(ONPStack)-1]
			case "+":
				result = ONPStack[len(ONPStack)-2] + ONPStack[len(ONPStack)-1]
			case "*":
				result = ONPStack[len(ONPStack)-2] * ONPStack[len(ONPStack)-1]
			case "/":
				result = ONPStack[len(ONPStack)-2] / ONPStack[len(ONPStack)-1]
			}
			//if token == "-" {
			//	result = ONPStack[len(ONPStack)-2] - ONPStack[len(ONPStack)-1]
			//} else if token == "+" {
			//	result = ONPStack[len(ONPStack)-2] + ONPStack[len(ONPStack)-1]
			//} else if token == "*" {
			//	result = ONPStack[len(ONPStack)-1] * ONPStack[len(ONPStack)-2]
			//} else if token == "/" {
			//	result = ONPStack[len(ONPStack)-2] / ONPStack[len(ONPStack)-1]
			//}
			ONPStack = ONPStack[:(len(ONPStack) - 2)]
			ONPStack = append(ONPStack, result)
		}
	}

	return ONPStack[len(ONPStack)-1]
}
