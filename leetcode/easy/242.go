package easy

/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	for _, v := range t {
		m[v]--
		if m[v] < 0 {
			return false
		}
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
