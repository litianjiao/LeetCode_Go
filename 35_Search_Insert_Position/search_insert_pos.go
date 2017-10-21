package lc0035

func SearchInsert(nums []int, target int) int {
	for i, val := range nums {
		if val >= target {
			return i
		}
	}
	//全部遍历完了也没找到只能插最后
	return len(nums)

}
