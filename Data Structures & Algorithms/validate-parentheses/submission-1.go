func isValid(s string) bool {
    stack := []rune{}
	for _, ch := range s {
		if len(stack) > 0 {
			if ch == '(' || ch == '[' || ch == '{' {
				stack = append(stack, ch)
			} else { 
				val := stack[len(stack) - 1]
				stack = stack[:len(stack) - 1]
				if ch == ')' && val != '(' {
					return false
				} else if ch == ']' && val != '[' {
					return false
				} else if ch == '}' && val != '{' {
					return false
				}
			}
		} else {
			stack = append(stack, ch)
		}
	}
	return len(stack) == 0
}
