package lc_0118

//递归求解，先确定两头的1，再确定中间的值
//返回的必须是一个二维数组（前N层）
func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	} else if numRows == 1 {
		return [][]int{[]int{1}}
	}

	child := generate(numRows - 1)
	cTail := child[len(child)-1] // child tail
	curr := make([]int, len(cTail)+1)
	curr[0] = 1
	curr[len(cTail)] = 1
	for i := 0; i < len(cTail)-1; i++ {
		curr[i+1] = cTail[i] + cTail[i+1]
	}
	return append(child, curr)
}
