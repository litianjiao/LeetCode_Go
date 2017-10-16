package NextGreaterElement2

//解法1:
//func NextGreaterElements(nums []int) []int {
//	lens := len(nums)
//	var array1 []int
//
//	for i := 0; i < lens; i++ {
//		array1 = append(array1, -1)
//
//		for j := 1; j < lens; j++ {
//			if nums[(i+j)%lens] > nums[i] {
//				array1[i] = nums[(i+j)%lens]
//				break
//			}
//		}
//	}
//	return array1
//}

//解法2：
func NextGreaterElements(nums []int) (res []int) {
	stack := make([]int, 0, len(nums))
	res = make([]int, len(nums))
	var lastMin int
	for i, n := range nums {
		if i == 0 || n <= lastMin {
			stack = append(stack, i)
			lastMin = n
		} else {
			j := len(stack) - 1
			for ; j >= 0; j-- {
				at := stack[j]
				if nums[at] < n {
					res[at] = n
				} else {
					break
				}
			}
			stack = append(stack[:j+1], i)
		}
	}
	if len(stack) > 0 {
		for i := 0; i < stack[0]; i++ {
			j := len(stack) - 1
			for ; j > 0; j-- {
				if nums[i] > nums[stack[j]] {
					res[stack[j]] = nums[i]
				} else {
					break
				}
			}
			stack = stack[:j+1]
		}
		for j := len(stack) - 1; j >= 0; j-- {
			if nums[stack[j]] < nums[stack[0]] {
				res[stack[j]] = nums[stack[0]]
			} else {
				res[stack[j]] = -1
			}
		}
	}
	return
}
