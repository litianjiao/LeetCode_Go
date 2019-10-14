
func isAnagram(s string, t string) bool {
	m1 := make([]int, 26)
	arr_s := []byte(s)
	arr_t := []byte(t)
	if len(arr_s) != len(arr_t) {
		return false
	}

	for i := 0; i < len(arr_s); i++ {
		m1[arr_s[i]-'a']++
	}
	for i := 0; i < len(arr_t); i++ {
		m1[arr_t[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if m1[i] != 0 {
			return false
		}
	}
	return true
}