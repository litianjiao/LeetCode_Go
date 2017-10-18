package roman2int

func RomanToInt(s string) int {
	//查表法
	var ret int
	m := map[rune]int{
		rune('I'): 1,
		rune('V'): 5,
		rune('X'): 10,
		rune('L'): 50,
		rune('C'): 100,
		rune('D'): 500,
		rune('M'): 1000,
	}
	//先拆分原string
	runes := []rune(s)
	for i := 0; i < len(runes)-1; i++ {
		if m[runes[i]] < m[runes[i+1]] {
			ret -= m[runes[i]]
		} else {
			ret += m[runes[i]]
		}
	}
	//补上最后一位
	ret += m[runes[len(runes)-1]]

	return ret
}
func RomanToInt2(s string) int {
	//查表法
	m := map[rune]int{
		rune('I'): 1,
		rune('V'): 5,
		rune('X'): 10,
		rune('L'): 50,
		rune('C'): 100,
		rune('D'): 500,
		rune('M'): 1000,
	}
	//先拆分原string
	runes := []rune(s)
	if len(runes) == 1 {
		return m[runes[0]]
	}
	ret := m[runes[0]]
	for i := 1; i < len(runes); i++ {
		if m[runes[i-1]] < m[runes[i]] {
			ret += m[runes[i]] - 2*m[runes[i-1]]
		} else {
			ret += m[runes[i]]
		}
	}
	return ret
}