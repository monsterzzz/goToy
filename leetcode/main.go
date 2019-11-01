package main

import (
	"fmt"
	"strings"
)

func main() {
	var s []string
	for _, v := range fmt.Sprintf("%b", uint32(4)) {
		s = append(s, string(v))
	}
	fmt.Println(strings.Join(s, ""))
}
