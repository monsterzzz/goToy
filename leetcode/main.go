package main

import "fmt"

func main() {
	nums := []int{3, 2, 4}
	target := 6
	m := make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}
	for k, v := range nums {
		if _, ok := m[target-v]; ok && m[target-v] != k {
			return []int{k, m[target-v]}
		}
	}
	_, ok := m[2]
	fmt.Println(ok)
}
