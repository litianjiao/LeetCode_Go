package lc0217

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

// para 是参数
type para struct {
	nums []int
}

// ans 是答案
type ans struct {
	one bool
}

func Test_containsDuplicate(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				[]int{1, 2, 3, 3},
			},
			ans{
				true,
			},
		},
		question{
			para{
				[]int{1, 2, 3, 4, 5},
			},
			ans{
				false,
			},
		},
		question{
			para{
				[]int{1, 2, 3, 4, 1},
			},
			ans{
				true,
			},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, containsDuplicate(p.nums), "输入:%v", p)
	}
}
