package lc0101

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return process(root.Left, root.Right)
}

func process(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	} else if l == nil || r == nil {
		/*防止其中一个为空*/
		return false
	}
	if l.Val != r.Val {
		return false
	}
	/*  // | \\      */
	return process(l.Left, r.Right) && process(l.Right, r.Left)
}
