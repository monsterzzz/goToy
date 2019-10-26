package main

import "fmt"

func main() {
	var a T
	var b = &a
	fmt.Println(&b)

	testA(b)
	fmt.Println(b)
	fmt.Println(&b)
	testB(&a)
	fmt.Println(a)
}

type T struct {
	name string
}

func testA(a *T) {
	fmt.Println(&a)
	//a = &T{name:"test"}
	a.name = "test"
	fmt.Println(&a)

}

func testB(a *T) {
	*a = T{"test"}
}
