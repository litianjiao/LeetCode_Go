package lc0069

/*
 *  这是求整数的开方结果，采用枚举法来寻找开方的值，为了最快效率,二分查找
 */
func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	l, h := 1, x
	for l <= h {
		mid := l + (h-l)/2
		res := x / mid
		if res == mid {
			return mid
		} else if res < mid {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	return h

}
