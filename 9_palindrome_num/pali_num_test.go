package lc_0009

import "testing"

func TestPaliInt(t *testing.T) {
	testInt := []int{12345, 123321, 12321, 100}
	res := []bool{false, true, true, false}
	caseNum := 4
	for i := 0; i < caseNum; i++ {
		if ret := IisPalindrome(testInt[i]); ret != res[i] {
			t.Fatalf("case %d fails:%v\n", i, ret)
		}
	}
}
