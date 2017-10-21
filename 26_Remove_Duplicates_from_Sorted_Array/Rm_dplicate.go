package lc_0026

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	index := 1
	for i := 1; i < len(nums); i++ {
		if nums[index-1] != nums[i] {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}
