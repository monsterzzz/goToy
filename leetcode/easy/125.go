package easy

import "fmt"

/*
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:

输入: "race a car"
输出: false


*/

func IsPalindrome(s string) bool {

	if len(s) == 0 {
		return true
	}

	p1 := 0
	p2 := len(s) - 1

	for p1 != len(s)/2 {
		fmt.Println(p1, p2, rune(s[p1]), rune(s[p2]))
		if !((65 <= rune(s[p1]) && rune(s[p1]) <= 65+26) || (97 <= rune(s[p1]) && rune(s[p1]) <= 97+26) || (48 <= rune(s[p1]) && rune(s[p1]) <= 57)) {
			p1++
			continue
		}
		if !((65 <= rune(s[p2]) && rune(s[p2]) <= 65+26) || (97 <= rune(s[p2]) && rune(s[p2]) <= 97+26) || (48 <= rune(s[p2]) && rune(s[p2]) <= 57)) {
			p2--
			continue
		}
		fmt.Println("??")
		var tmp1 int
		var tmp2 int
		if 65 <= rune(s[p1]) && rune(s[p1]) <= 65+26 {
			tmp1 = int(s[p1]) + 32
		} else {
			tmp1 = int(s[p1])
		}

		if 65 <= rune(s[p2]) && rune(s[p2]) <= 65+26 {
			tmp2 = int(s[p2]) + 32
		} else {
			tmp2 = int(s[p2])
		}

		fmt.Println(tmp1, tmp2, s[p1], s[p2], string(tmp1), string(tmp2))

		if tmp1 != tmp2 {
			return false
		}

		p1++
		p2--

	}
	return true

}
