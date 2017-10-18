package longest_com_pre

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	testStr := [][]string{
		[]string{"ss111", "sss111", "s111", "aaaaaa"},
		[]string{"ss111", "a111", "s111", "aaaaaa"},
		[]string{"ss111", "ss111", "ss111", "ss111"},
	}
	res := []string{
		"",
		"",
		"ss111",
	}

	caseNum := 3
	for i := 0; i < caseNum; i++ {
		if ret := LongestCommonPrefix(testStr[i]); ret != res[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
