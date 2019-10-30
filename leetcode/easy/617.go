package easy

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil { // t1 为空 则返回 t2 // 已经包含 t1 t2 都为空的情况
		return t2
	}
	if t2 == nil { // t2 为空 则返回 t1
		return t1
	}

	t1.Val += t2.Val // t1 t2 不为空 则相加
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}
