package easy

import "fmt"

/*
给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。
*/

func repeatedSubstringPattern(s string) bool {
	if len(s) <= 1 {
		return false
	}

	tmp := s + s // abaaba
	fmt.Println(tmp, len(tmp))
	tmp = tmp[1 : len(tmp)-1] // baab
	fmt.Println(tmp)
	n := 0
	for {

		if tmp[n:n+len(s)] == s { // b
			return true
		}
		n++
		if n+len(s) > len(tmp) {
			break
		}
	}

	return false

}
