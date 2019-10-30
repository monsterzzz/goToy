package main

import (
	"fmt"
	"goToy/leetcode/easy"
)

func main() {
	var a []int
	a = append(a, 8, 8, 7, 7, 7)
	fmt.Println(easy.MajorityElement(a))
}
