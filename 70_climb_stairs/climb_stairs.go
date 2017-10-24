package lc_0070


//fibonacci
func climbStairs(n int) int {

	if n<=0 {return 0}
	if n==1 {return 1}
	if n==2{return 2}

	step := make([]int, n+1)
	step[0], step[1] = 1, 1

	for i := 2; i <= n; i++ {
		step[i] = step[i-1] + step[i-2]
	}

	return step[n]
}

