package easy

/*
4.
请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点，你将只被给定要求被删除的节点。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	//  无法得到前置节点，使用替身
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
