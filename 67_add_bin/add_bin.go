package lc0067

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	ls := len(a)
	binA := trans(a, ls)
	binB := trans(b, ls)

	return binToStr(add(binA, binB))
}

//转换成二进制并翻转每位的顺序
func trans(s string, lens int) []int {
	res := make([]int, lens)
	ls := len(s)
	for i, v := range s {
		res[lens-ls+i] = int(v - '0')
	}
	return res
}

func add(a, b []int) []int {
	lens := len(a) + 1
	res := make([]int, lens)
	for i := lens - 1; i >= 1; i-- {
		temp := res[i] + a[i-1] + b[i-1]
		res[i] = temp % 2
		res[i-1] = temp / 2
	}
	i := 0
	//还原顺序
	for i < lens-1 && res[i] == 0 {
		i++
	}
	return res[i:]
}

func binToStr(bin []int) string {
	bytes := make([]byte, len(bin))
	for i := range bytes {
		bytes[i] = byte(bin[i]) + '0'
	}

	return string(bytes)
}
