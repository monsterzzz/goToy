package easy

/*
有效括号字符串为空 ("")、"(" + A + ")" 或 A + B，其中 A 和 B 都是有效的括号字符串，+ 代表字符串的连接。例如，""，"()"，"(())()" 和 "(()(()))" 都是有效的括号字符串。

如果有效字符串 S 非空，且不存在将其拆分为 S = A+B 的方法，我们称其为原语（primitive），其中 A 和 B 都是非空有效括号字符串。

给出一个非空有效字符串 S，考虑将其进行原语化分解，使得：S = P_1 + P_2 + ... + P_k，其中 P_i 是有效括号字符串原语。

对 S 进行原语化分解，删除分解中每个原语字符串的最外层括号，返回 S 。

*/

func removeOuterParentheses(S string) string {
	stack := make([]rune, 0)
	var res string
	start := 0
	for i, v := range S {
		if string(v) == "(" { // 当遇到左括号就入列
			stack = append(stack, v)
		} else {
			stack = stack[:len(stack)-1] // 遇到右括号就出列
			if len(stack) == 0 {         // 如果当前列为空，就说明找到了一组(()())
				res += S[start+1 : i] // 当前位置为 最后的 ) ,截取中间字符串
				start = i + 1         // 设定下次开始的位置为当前的 ) 的下一位字符
			}
		}
	}
	return res
}
