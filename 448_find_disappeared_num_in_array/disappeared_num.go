package lc_0448

/*
******************************************************************
  * @brief  解法1 遍历两次该切片，第一次将存在数下标中的数字置为负数，第二次直接找val是正数的
  *				index,存入res
  * @param
  * @ret
  * @author    Troy
  * @date      2018/2/6 0:46
******************************************************************
*/
func findDisappearedNumbers(nums []int) []int {
	var res []int
	for _, num := range nums {
		index := abs(num) - 1
		nums[index] = -abs(nums[index])
	}
	for idx, num := range nums {
		if num > 0 {
			res = append(res, idx+1)
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/*
******************************************************************
  * @brief  解法2
  * @param
  * @ret
  * @author    Troy
  * @date      2018/2/6 0:46
******************************************************************
*/
func findDisappearedNumbers2(nums []int) []int {

	ans := []int{}
	for i, _ := range nums {
		for i != nums[i]-1 {
			if nums[nums[i]-1] == nums[i] {
				break
			}
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i, _ := range nums {
		if nums[i]-1 != i {
			ans = append(ans, i+1)
		}
	}

	return ans
}
