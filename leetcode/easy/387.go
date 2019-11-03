package easy

/*
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

案例:

s = "leetcode"
返回 0.

s = "loveleetcode",
返回 2.


*/

func firstUniqChar(s string) int {
	m := make(map[rune]int)

	for _, v := range s {
		m[v]++
	}

	for i, v := range s {
		if m[v] == 1 {
			return i + 1
		}
	}
	return -1

}
