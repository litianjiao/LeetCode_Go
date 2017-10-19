package lc_0005
//解法 双指针扫描
func longestPalindrome(s string) string {
   //肯定是回文
	if len(s) == 0 {
		return s
	}

	length := len(s)
	// 最长回文的首字符索引，和最长回文的长度
	start, max := 0, 1
	// l 代表回文`正中间段`首字符的索引号
	for i := 1; i < length; i++ {
		l := i - 1
		for l >= 0 && s[l] == s[i] {
			l--
		}
    //r代表回文末字符的索引号
		r := i + 1
		for r < length && s[r] == s[i] {
			r++
		}

		i = r - 1
      //判断此次缩到的是否是回文
		for l >= 0 && r < length {
			if s[l] != s[r] {
				break
			}
			l--
			r++
		}

		size := r - l - 1
		//如果创新记录则更新记录
		if size > max {
			max = size
			start = l + 1
		}
	}

	return s[start:(start + max)]
}

/*
******************************************************************
  * @brief     is Palindrome or Not
  * @param
  * @ret       bool is Palindrome or not
  * @author    Troy
  * @date      2017/10/19 20:42
******************************************************************
*/
func findPalindrome(subRune []rune, i, j int) bool {
for i<j{
	if subRune[i]!=subRune[j]{
		return false
	}
	i++
	j--
}
return true

}
