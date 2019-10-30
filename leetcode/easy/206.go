package easy

/*
反转一个单链表。
*/

func ReverseList(head *ListNode) *ListNode {
	var pre, cur, next *ListNode
	cur = head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur

		cur = next
	}
	return pre
}

// 1->2->3->4->5

// 1 -> 2 -> 3
// tmpNode : 2 -> 1
// head : 2 -> 3
