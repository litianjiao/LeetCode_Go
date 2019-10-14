package lc0038

import (
	"bytes"
)

/* n=1时输出字符串1；n=2时，数上次字符串中的数值个数，因为上次字符串有1个1，所以输出11；n=3时，
 * 由于上次字符是11，有2个1，所以输出21；n=4时，由于上次字符串是21，有1个2和1个1，所以输出1211。
 * 依次类推，写个countAndSay(n)函数返回字符串。
 */

func countAndSay(n int) string {
	s := "1"
	for i := 1; i < n; i++ {
		s = strCountAndSay(s)
	}
	return s
}

func strCountAndSay(s string) string {
	buf := bytes.NewBuffer([]byte{})

	prev := []rune(s)[0]
	count := 0

	write := func() {
		buf.WriteRune(rune(count) + '0')
		buf.WriteRune(prev)
	}

	for _, c := range s {
		if c == prev {
			count++
		} else {
			write()
			prev = c
			count = 1
		}
	}
	write()
	return buf.String()
}

// func countAndSay(n int) string {
// 	if n <= 0 {
// 		return ""
// 	} else if n == 1 {
// 		return "1"
// 	}
// 	rets := countAndSayInts(n)
// 	bytes := make([]byte, len(rets))
// 	intTobyte := []byte("123456789")
// 	for index, ret := range rets {
// 		bytes[index] = intTobyte[ret-1]
// 	}
// 	return string(bytes)
// }

// func countAndSayInts(n int) (rets []int) {
// 	if n < 1 {
// 		return []int{}
// 	} else if n == 1 {
// 		return []int{1}
// 	}
// 	rets = []int{1}
// 	for i := 2; i <= n; i++ {
// 		tmps := []int{}
// 		pre, cnt := -1, 0
// 		for index, val := range rets {
// 			if index == 0 {
// 				pre = val
// 				cnt = 1
// 				continue
// 			}
// 			if val != pre {
// 				tmps = append(tmps, cnt, pre)
// 				pre = val
// 				cnt = 1
// 			} else {
// 				cnt++
// 			}
// 		}
// 		tmps = append(tmps, cnt, pre)
// 		rets = tmps
// 	}
// 	return
// }
