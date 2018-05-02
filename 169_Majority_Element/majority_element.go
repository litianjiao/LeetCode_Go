package lc0169

func majorityElement(nums []int) int {
	//len >0  该数字出现次数>n/2
	count := 0
	var data int
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			data = nums[i]
		}

		if nums[i] == data {
			count += 1
		} else {
			count += -1
		}
	}
	return data
}
