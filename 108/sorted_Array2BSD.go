package lc0108

/**
 * Definition for a binary tree node.
 *
 * */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return fn(&nums, 0, len(nums)-1)
}

func fn(nums *[]int, start int, end int) *TreeNode {
	if start > end {
		return nil
	}
	m := start + (end-start)/2
	return &TreeNode{
		(*nums)[m],
		fn(nums, start, m-1),
		fn(nums, m+1, end),
	}
}
