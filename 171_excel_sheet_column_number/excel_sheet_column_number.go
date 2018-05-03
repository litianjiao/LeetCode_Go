package lc0171

func titleToNumber(s string) int {
	sum := 0
	for _, ch := range s {
		current := int(ch-'A') + 1
		//N+N+N  从前往后每一位相当于26个
		sum = sum*26 + current
	}
	return sum
}
