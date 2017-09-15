package singleNumber

//136 给定一个整数数组，里面只有一个数只出现一次，其余的数都出现两次，找出只出现一次的数。
//所有元素异或，因为   a^a=0  a^0=a  出现两次的数会相消，留下只出现一次的数
func singleNumber(nums []int) int {
	var ret int
	for _, v := range nums {
		ret ^= v
	}
	return ret
}
