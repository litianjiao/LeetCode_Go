package lc0198
//最值方案问题考虑采用动态规划

func rob(nums []int) int {
    size:=len(nums)
    if size==0{
        return 0
    }
    if size==1{
        return nums[0]
    }
    dp:=nums[:]//dp表示第N个房子时能偷到的最大值
    dp[0]=nums[0]
    dp[1]=max(nums[0],nums[1])
    
    //在这里比较现在这个偷还是不偷
    for i:=2;i<size;i++{
    dp[i]=max(dp[i-1],nums[i]+dp[i-2])
    }
    return dp[size-1]
}

func max(a,b int)int{
    if a>b{return a}
    return b
}


//节省空间，直接倒数两个
func rob2(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) < 2 {
        return nums[0]
	}
	
	t2 := 0  //如果当前要偷，则总额
	t1 := 0  //如果当前不偷，则总额

    for i := 0; i < len(nums); i++ {
		t := t2 + nums[i]
		if t < t1 {
			t = t1
		}

		t2 = t1
		t1 = t
	}	
	return t1
}