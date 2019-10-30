package main

import (
	"fmt"
	"goToy/leetcode/easy"
)

func main() {

	a := []int{0, 1, 0, 3, 12}
	easy.MoveZeroes(a)
	fmt.Println(a)
}
