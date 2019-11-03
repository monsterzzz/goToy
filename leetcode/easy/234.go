package easy

/*
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func IsPalindrome1(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}

	// 找到中点
	fast := head.Next.Next
	slow := head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// 反转后半部分
	if fast != nil {
		slow = slow.Next
	}
	//var pre *ListNode
	var pre *ListNode

	for slow != nil {
		tmp := slow.Next
		slow.Next = pre
		pre = slow
		slow = tmp
	}

	for pre != nil {
		if head.Val != pre.Val {
			return false
		}
		head = head.Next
		pre = pre.Next
	}
	return true

}
