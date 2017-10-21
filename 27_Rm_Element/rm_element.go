package lc_0027

func RemoveElement(nums []int, val int) int {
	j:=0
	for i:=0;i<len(nums);i++ {
		if nums[i]==val {
			continue
		}
		nums[j]=nums[i]
		j++
	}
	  return j

}