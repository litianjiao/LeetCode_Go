package lc_0448

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

type para struct {
	one []int
}

type ans struct {
	one []int
}

func Test_Problem0448(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[]int{4, 3, 2, 7, 8, 2, 3, 1}},
			ans{[]int{5, 6}},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, findDisappearedNumbers2(p.one), "输入:%v", p)
	}
}
