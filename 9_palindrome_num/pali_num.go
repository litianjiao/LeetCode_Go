package lc_0009

import (
	"strconv"
)

/*
******************************************************************
  * @brief     回文 字符串解法  233333
  * @param
  * @ret
  * @author    TRoy
  * @date      2017/9/17 22:59
******************************************************************
*/
func AisPalindrome(x int) bool {

	strX := strconv.Itoa(x)
	t := []rune(strX)
	len := len(t)

	for i, _ := range t {
		if i == len/2 {
			break
		}
		last := len - i - 1
		if t[i] != t[last] {
			return false
		}

	}
	return true
}

/*
******************************************************************
  * @brief     如果回文是纯数字解法  非常math
  * @param
  * @ret
  * @author    TRoy
  * @date      2017/9/17 22:59
******************************************************************
*/
func IisPalindrome(x int) bool {
	left, right := x, 0
	for left > 0 {
		right = left%10 + right*10
		left = left / 10
	}
	return x == right
}
