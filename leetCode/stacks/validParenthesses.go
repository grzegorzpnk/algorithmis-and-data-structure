package stacks

func IsValid(s string) bool {

	stack := make([]rune, 0)

	for _, char := range s {

		if len(stack) == 0 {
			stack = append(stack, char)
			continue
		}
		if (char == '}' && stack[len(stack)-1] == '{') || (char == ']' && stack[len(stack)-1] == '[') || (char == ')' && stack[len(stack)-1] == '(') {
			stack = stack[:len(stack)-1]
			return false
		} else {
			stack = append(stack, char)
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}
