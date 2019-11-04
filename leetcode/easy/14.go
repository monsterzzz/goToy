package easy

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。
*/

func LongestCommonPrefix(strs []string) string {
	ans := strs[0]

	for i := 1; i < len(strs); i++ {
		var tmp string
		for j, v := range strs[i] {
			if j >= len(ans) {
				break
			}

			if strs[i][j] == ans[j] {
				tmp += string(v)
			} else {
				break
			}
		}
		if tmp != "" {
			ans = tmp
		} else {
			return ""
		}
	}
	return ans
}
