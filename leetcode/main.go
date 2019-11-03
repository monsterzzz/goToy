package main

import (
	"fmt"
	"goToy/leetcode/easy"
)

func main() {

	//l1 := &easy.ListNode{Val:1}
	//l2 := &easy.ListNode{Val:2}
	//l3 := &easy.ListNode{Val:3}
	//l4 := &easy.ListNode{Val:4}
	//l5 := &easy.ListNode{Val:5}
	//l1.Next = l2
	//l2.Next = l3
	//l3.Next = l4
	//l4.Next = l5

	l1 := &easy.ListNode{Val: 1}
	l2 := &easy.ListNode{Val: 2}
	l3 := &easy.ListNode{Val: 2}
	l4 := &easy.ListNode{Val: 1}

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	fmt.Println(easy.IsPalindrome1(l1))

}
