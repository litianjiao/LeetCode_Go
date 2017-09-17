package atoi

import (
	//"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

func myAtoi(str string) int {
	trimed := strings.TrimSpace(str)
	nonNumric := len(trimed)
	for i, rune := range trimed {
		//fmt.Printf("trimed[%d] is rune: %s\n", i, rune)
		if !unicode.IsDigit(rune) && rune != '-' && rune != '+' {
			nonNumric = i
			break
		}
	}
	sanitized := trimed[:nonNumric]
	res, _ := strconv.Atoi(sanitized)
	if res >= math.MaxInt32 {
		res = math.MaxInt32
	}
	if res < math.MinInt32 {
		res = math.MaxInt32
	}
	return res
}
