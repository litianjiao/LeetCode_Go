package lc0058

func lengthOfLastWord(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	ret := 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if ret != 0 {
				return ret
			}
			continue
		}
		ret++
	}

	return ret
}
