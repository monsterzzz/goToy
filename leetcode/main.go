package main

import "goToy/leetcode/easy"

func main() {
	n1 := easy.ListNode{
		Val: 1,
		Next: &easy.ListNode{
			Val: 2,
			Next: &easy.ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	n2 := easy.ListNode{
		Val: 1,
		Next: &easy.ListNode{
			Val: 3,
			Next: &easy.ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	easy.MergeTwoLists(&n1, &n2)
}
