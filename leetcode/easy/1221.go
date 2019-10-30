package easy

/*
在一个「平衡字符串」中，'L' 和 'R' 字符的数量是相同的。

给出一个平衡字符串 s，请你将它分割成尽可能多的平衡字符串。

返回可以通过分割得到的平衡字符串的最大数量。

*/

func balancedStringSplit(s string) int {
	num := 0
	res := 0
	for _, r := range s {
		if string(r) == "L" {
			num++
		} else {
			num--
		}
		if num == 0 {
			res++
		}

	}
	return res
}
