package main

import "fmt"

func main() {
	//doutu.Run()
	sum(&calc{name: "a"})
}

type a interface {
	test(a, b int) int
}

type calc struct {
	name string
}

func (c *calc) test(a, b int) int {
	return a + b
}

func sum(c a) {
	fmt.Println(c.test(1, 1))
}
