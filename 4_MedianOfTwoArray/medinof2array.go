package lc_0004

// Median of Two Sorted Arrays
//俩数组先合并，再求解中位数
//解法 合并过程非常重要，这里是交错扫描，在arr1中找arr2插入的位置，在arr2中找arr1能插入的位置，因为原数组已经是排好序的了
//那么什么条件下能满足另外一个数组插到合适的位置呢
//保证每次都是一人只走一步，除非走完
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := combine(nums1, nums2)
	return medianOf(nums)

}

func combine(arr1 []int, arr2 []int) []int {
	idx1, idx2, len1, len2 := 0, 0, len(arr1), len(arr2)
	res := make([]int, len1+len2)
	for k := 0; k < len1+len2; k++ {
		if idx1 == len1 ||
			(idx1 < len1 && idx2 < len2 && arr1[idx1] > arr2[idx2]) {
			res[k] = arr2[idx2]
			idx2++
			continue
		}
		if idx2 == len2 ||
			(idx1 < len2 && arr1[idx1] <= arr2[idx2]) {
			res[k] = arr1[idx1]
			idx1++
		}
	}
	return res
}

func medianOf(nums []int) float64 {
	len := len(nums)
	if len == 0 {
		panic("空切片")
	}

	if len%2 == 0 {
		return float64(nums[len/2]+nums[len/2-1]) / 2.0
	}

	return float64(nums[len/2])
}
