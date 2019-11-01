package easy

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(leftNode, rightNode *TreeNode) bool {
	if leftNode == nil && nil == rightNode {
		return true
	}
	if leftNode == nil || rightNode == nil {
		return false
	}
	return (leftNode.Val == rightNode.Val) && isMirror(leftNode.Left, rightNode.Right) && isMirror(leftNode.Right, rightNode.Left)
}
