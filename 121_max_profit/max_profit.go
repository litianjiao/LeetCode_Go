package lc0121

//通过每个节点之间的差值来看盈亏，步步都更赚

func maxProfit(prices []int) int {
	lenA := len(prices)
	var ret, maxCurrent int
	for i := 1; i < lenA; i++ {
		maxCurrent += prices[i] - prices[i-1]
		maxCurrent = max(0, maxCurrent)
		ret = max(ret, maxCurrent)
	}
	return ret
} 



func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

