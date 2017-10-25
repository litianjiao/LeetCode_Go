package lc_0020

//用栈的方法解决
func isValid(s string) bool {

	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		//push
		case '(': stack = append(stack, ')')
		case '{': stack = append(stack, '}')
		case '[': stack = append(stack, ']')
		default:
			size := len(stack)
			if size > 0 && c == stack[size - 1] {
				//pop 只要前N个
				stack = stack[:size - 1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
