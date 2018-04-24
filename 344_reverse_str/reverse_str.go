package lc0344

import "fmt"

//之前写法
func reverseString2(s string) string {
	var result string
	strlen := len(s)
	for i := 0; i < strlen; i++ {
		result = result + fmt.Sprintf("%c", s[strlen-1-i])
	}
	return result
}

//优化写法 先强制类型转换为byte 再通过数组索引直接交换
func reverseString(s string) string {
	bytes := []byte(s)
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		tmp := bytes[i]
		bytes[i] = bytes[j]
		bytes[j] = tmp
	}
	return string(bytes)
}
