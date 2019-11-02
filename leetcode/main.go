package main

import (
	"goToy/leetcode/easy"
)

func main() {

	head := &easy.ListNode{Val: 3}
	l1 := &easy.ListNode{Val: 2}
	l2 := &easy.ListNode{Val: 0}
	l3 := &easy.ListNode{Val: -4}

	head.Next = l1
	l1.Next = l2
	l2.Next = l3
	l3.Next = l1

}
