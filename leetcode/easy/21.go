package easy

/*
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
*/

func MergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	cur1 := l1
	cur2 := l2
	var n ListNode
	res := &n
	for cur1 != nil && cur2 != nil {
		if cur1.Val > cur2.Val {
			res.Next, cur2 = cur2, cur2.Next
		} else {
			res.Next, cur1 = cur1, cur1.Next
		}
		res = res.Next
	}
	if cur1 != nil {
		res.Next = cur1
	} else {
		res.Next = cur2
	}
	return res.Next
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l2.Next, l1)
		return l2
	}

}
