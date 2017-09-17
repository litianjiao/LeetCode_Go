package atoi

import "testing"

func TestAtoi(t *testing.T) {
	testStr := []string{"200", "20000000", "-4444", "-66666666"}
	retInt := []int{200, 20000000, -4444, -66666666}
	caseNum := 4
	for i := 0; i < caseNum; i++ {
		if ret := myAtoi(testStr[i]); ret != retInt[i] {
			t.Fatalf("case %d fails:%v\n", i, ret)
		}
	}
}
