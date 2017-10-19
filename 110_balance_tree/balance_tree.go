package lc_0110

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
//DP解
func balance(root *TreeNode) (int, bool) {

	if root == nil {
		return 1, true
	}
	//递归分别查找
	ld, ok := balance(root.Left)
	if !ok {
		return -1, false
	}

	rd, ok := balance(root.Right)
	if !ok {
		return -1, false
	}
	//diff是平衡因子，平衡因子取值只能是0 1 -1
	//如若该节点的子节点左右高度差的绝对值已失衡，则直接可判断此树不平衡
	diff := ld - rd
	if diff > 1 || diff < -1 {
		return -1, false
	}
	//树的深度就是其左、右子树深度的较大值再加1
	if diff > 0 {
		return ld + 1, true
	}
	return rd + 1, true
}

func isBalanced(root *TreeNode) bool {
	//可以是空树
	if root == nil {
		return true
	}
	//先判断根节点的左右子节点的高度差是否不大于1，然后递归判断子节点，如果有高度差大于1的节点，则该树不是平衡二叉树。
	ld, ok := balance(root.Left)
	if !ok {
		return false
	}

	rd, ok := balance(root.Right)
	if !ok {
		return false
	}

	diff := ld - rd
	if diff > 1 || diff < -1 {
		return false
	}
	return true
}
