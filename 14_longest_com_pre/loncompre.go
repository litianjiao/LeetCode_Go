package lc_0014

import "strings"

//最长公共前缀 用指针解
func LongestCommonPrefix(strs []string) string {
	//size表示有几个字符串需要找最长公共前缀
	size := len(strs)
	if size == 0 {
		return ""
	}
	if size == 1 {
		return strs[0]
	}
	//初始值为错误值""
	prefix := ""
	//各个字符串元素逐字符判断 i是待判断数组下标，只有在所有字符串都满足时才+1
	//  j是数组内字节下标
	for i, j := 0, len(strs[0]); i < j && j > 0; i++ {
		//在每个之后的字符串内找第一个字符串逐个前缀
		prefix = strs[0][0 : i+1]
		for k := 1; k < size; k++ {
			if !strings.HasPrefix(strs[k], prefix) {
				return strs[0][0:i]
			}
		}
	}
	//这几个都长得一样
	return prefix
}

// Go原生实现hasprefix
// len(s) >= len(prefix) && s[0:len(prefix)] == prefix
