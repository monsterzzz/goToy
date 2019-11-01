package easy

/*
编写一个程序，找到两个单链表相交的起始节点。
*/

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	startA, startB := headA, headB
	for headA != headB { // 如果不相交 必然会有一个时刻都指向nil 如果相交必然有一个时刻相等

		if headA == nil { // 如果达到了最后一个节点
			headA = startB //  指向 链表B
		} else {
			headA = headA.Next
		}

		if headB == nil {
			headB = startA
		} else {
			headB = headB.Next
		}

	}
	return headA
}
