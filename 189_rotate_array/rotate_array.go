package lc0189

func rotate(nums []int, k int) {
	idx := len(nums) - k%len(nums) //取余，Nl==0

	//因为要求使用原地算法，所以不另外开辟空间做缓存，直接通过交换来移动

	reverse(nums[0:idx])
	reverse(nums[idx:])
	reverse(nums)
}

func reverse(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}
