package lc_0033

func search(nums []int, target int) int {
	lo := 0
	len := len(nums)
	hi := len - 1
	//首先确定旋转的索引在mid的左边还是右边
	//如果左边元素小于mid，那么mid一定是落在了右面的递增序列中；反之一定是落在了左边
	var mid int
	for lo < hi {
		mid = (hi + lo) / 2
		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	//得到了新的数列后可以判断target是否在数组中了
	var rot int
	rot = lo
	lo = 0
	hi = len - 1
	for lo <= hi {
		mid = (lo + hi) / 2
		realmid := (mid + rot) % len
		if nums[realmid] == target {
			return realmid
		}
		if nums[realmid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}
