package lc_0001
//HASH map solution
func twoSum(nums []int, target int) []int {
	// m := make(map[int]int)
	// for i := 0; i < len(nums); i++ {
	// 	another := target - nums[i]
	// 	if _, ok := m[another]; ok { //2.如果序号所对应的键值存在则说明有
	// 		return []int{m[another], i}
	// 	}
	// 	m[nums[i]] = i //1.通过将键值与序号对应
	// }
	// return nil
//核心步骤是在hashmap中找target-num 是否存在而且这个num不在其本身
	list := make(map[int]int, len(nums))
	for index, num := range nums {
		if index_2, isComplex := list[target-num]; isComplex && index != index_2 {
			return []int{index_2, index}
		} else {
			list[num] = index
		}
	}
	return []int{0, 0}
}
