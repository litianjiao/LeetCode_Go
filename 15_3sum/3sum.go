package threesum

import "sort"

//解法1
func threeSum1(nums []int) (sums [][]int) {
	sort.Ints(nums)
	len := len(nums)
	for i := 0; i < len-1; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := 0 - nums[i]
		j := i + 1
		k := len - 1
		for j < k {
			if target == nums[j]+nums[k] {
				sums = append(sums, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			} else if nums[j]+nums[k] > target {
				k--
			} else {
				j++
			}
		}
	}
	return sums
}

//解法1 优化版  beats 100% golang =.=
func threeSum(nums []int) (sums [][]int) {
	sort.Ints(nums)
	l := len(nums)
	//i表示所有元素都有可能作为第一个元素而进行遍历
	for i := 0; i < l-1; i++ {
		idx1 := i + 1
		idx2 := l - 1
		for idx1 < idx2 {

			if nums[idx1]+nums[idx2]+nums[i] < 0 {
				//值还不够 index1得向后推进
				idx1++
			} else if nums[idx1]+nums[idx2]+nums[i] > 0 {
				//值太大了 index2向前推进
				idx2--
			} else {
				//和恰好为0
				sums = append(sums, []int{nums[i], nums[idx1], nums[idx2]})
				//因为排序过，所以两边一同缩进index，跳过相同的值
				for idx1 < idx2 && nums[idx1] == nums[idx1+1] {
					idx1++
				}

				for idx2 > idx1 && nums[idx2] == nums[idx2-1] {
					idx2--
				}
				idx1++
				idx2--
			}
		}
		//跳过相等的元素
		for ; i < l-1 && nums[i] == nums[i+1]; i++ {
		}
	}
	return sums
}
