package main

import (
	"fmt"
)

func main() {
	n := 3
	initQue := []int{1, 2}
	for i := 2; i < n; i++ {
		initQue = append(initQue, initQue[i-1]+initQue[i-2])
	}
	fmt.Println(initQue[n-1])
}
