package lc_0503

import "testing"

func TestNextGreaterElements(t *testing.T) {
	testInt := [][]int{
		[]int{1, 2, 1},
		[]int{4, 2, 1},
		[]int{4, 2, 1, 4, 5, 6},
	}
	res := [][]int{
		[]int{2, -1, 2},
		[]int{-1, 4, 4},
		[]int{5, 4, 4, 5, 6, -1},
	}

	caseNum := 1
	for i := 0; i < caseNum; i++ {
		if ret := NextGreaterElements(testInt[i]); ret[0] != res[i][0] && ret[1] != res[i][1] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}

}
