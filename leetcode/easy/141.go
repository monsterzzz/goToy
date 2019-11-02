package easy

/*
给定一个链表，判断链表中是否有环。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {

	cur1 := head
	cur2 := head

	for {
		if cur1 != nil {
			cur1 = cur1.Next
		}
		if cur2 != nil {
			cur2 = cur2.Next
			if cur2 != nil {
				cur2 = cur2.Next
			}
		}

		if cur2 == cur1 {
			break
		}

	}
	return (cur1 == cur2) && (cur1 != nil)
}
