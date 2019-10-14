package lc0125

import "strings"

func isPalindrome(s string) bool {
	r := []rune(strings.ToLower(s))
	rlen := len(r)
	for i, j := 0, rlen-1; i < j; {
		if !((r[i] >= 'a' && r[i] <= 'z') || (r[i] >= '0' && r[i] <= '9')) {
			i++
			continue
		}
		if !((r[j] >= 'a' && r[j] <= 'z') || (r[j] >= '0' && r[j] <= '9')) {
			j--
			continue
		}
		if r[i] != r[j] {
			return false
		}
		i++
		j--
	}
	return true
}
