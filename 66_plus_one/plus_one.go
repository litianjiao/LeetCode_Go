package lc_0066

/*大数相加*/
func plusOne(digits []int) []int {
	n := len(digits)
	if n == 0 {
		return []int{}
	}
	digits[n-1] = digits[n-1] + 1
	for i := n - 1; i >= 0; i-- {
		if digits[i] == 10 {
			digits[i] = 0
			if i == 0 {
				return append([]int{1}, digits...)
			}
			digits[i-1] = digits[i-1] + 1
		} else {
			return digits
		}
	}

	return digits
}
