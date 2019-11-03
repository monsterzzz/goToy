package main

import (
	"fmt"
)

func main() {

	s := "{[]}"

	m := make(map[rune]rune)
	m['('] = ')'
	m['{'] = '}'
	m['['] = ']'
	l := make([]rune, 0)
	for _, v := range s {
		fmt.Println(l)
		if len(l) == 0 {
			l = append(l, v)
			continue
		}

		if _, ok := m[l[len(l)-1]]; !ok {
			return false
		}

		if m[l[len(l)-1]] == v {
			//fmt.Println("!",l,string(l[len(l)-1]),string(m[l[len(l)-1]]),string(v))
			l = l[:len(l)-1]
		} else {
			l = append(l, v)
		}
	}

}
