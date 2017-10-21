package lc_0053

func getMax(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
******************************************************************
  * @brief  最大和
  * @param
  * @ret
  * @author    Troy
  * @date      2017/10/22 2:16
******************************************************************
*/
//解法 DP，
// 设dp[i]是数组a [0, i]区间最大的值，那么dp[i + 1] = max(dp[i], dp[i] + a[i + 1])
//不断合并已经获得最大值组合的子和
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max:=nums[0]
	dp:=nums[0]

	for i := 1; i < len(nums); i++ {
		dp=getMax(dp+nums[i],nums[i])
		max=getMax(dp,max)
	}
return max

}

