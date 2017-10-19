package lc_0007

import "math"

func reverse(x int) int {
	var right int64
	for x != 0 {
		right *= 10
		right = right + int64(x)%10
		//规定溢出时输出0
		if right > math.MaxInt32 || right < math.MinInt32 {
			return 0
		}

		x = x / 10
	}
	return int(right)

}
