package lc_0028

func strStr(haystack string, needle string) int {
	if haystack == needle {
		return 0
	}
	//偏移逐位置比对
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
