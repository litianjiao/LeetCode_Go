package lc0217

func containsDuplicate(nums []int) bool {
	//用map记录值
	m := make(map[int]bool, len(nums))
	for _, i := range nums {
		if m[i] {
			return true
		}
		m[i] = true
	}
	return false
}
