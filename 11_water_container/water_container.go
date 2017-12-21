package lc0011

/*一个容器内能装多少水是由其木板最低处所决定的*/
func maxArea(height []int) int {
	n := len(height)
	lo, hi := 0, n-1
	var res int

	for lo < hi {
		area := (hi - lo) * min(height[lo], height[hi])
		res = max(area, res)
		if height[lo] < height[hi] {
			lo++
		} else {
			hi--
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
