package lc0172

func trailingZeroes(n int) int {
	//跟着几个0<=因数中2*5的个数，因为有5的阶乘，总是有2，所以我们只需计算1~n中有几个5因子就可以了
	//	return 0 if n == 0 else n / 5 + self.trailingZeroes(n / 5)
	// if n == 0 {
	// 	return 0
	// } else {
	// 	return n/5 + trailingZeroes(n/5)
	// }
	res := 0

	for n >= 5 {
		n /= 5
		res += n
	}

	return res
}
