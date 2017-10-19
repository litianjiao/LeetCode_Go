package lc_0013

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	testStr := []string{"DCXX", "DCXXI", "DCXXII", "CDCXXII"}
	res := []int{620, 621, 622, 522}
	caseNum := 4
	for i := 0; i < caseNum; i++ {
		if ret := RomanToInt(testStr[i]); ret != res[i] {
			t.Fatalf("way1 case %d fails:%v\n", i, ret)
		}
	}

}

func TestRomanToInt2(t *testing.T) {
	testStr := []string{"DCXX", "DCXXI", "DCXXII", "CDCXXII"}
	res := []int{620, 621, 622, 522}
	caseNum := 4
	for i := 0; i < caseNum; i++ {
		if ret := RomanToInt(testStr[i]); ret != res[i] {
			t.Fatalf("way2 case %d fails:%v\n", i, ret)
		}
	}

}
