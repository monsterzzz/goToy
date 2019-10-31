package easy

import (
	"fmt"
)

/*
报数序列是一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：

1.     1
2.     11
3.     21
4.     1211
5.     111221
1 被读作  "one 1"  ("一个一") , 即 11。
11 被读作 "two 1s" ("两个一"）, 即 21。
21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。

给定一个正整数 n（1 ≤ n ≤ 30），输出报数序列的第 n 项。

注意：整数顺序将表示为一个字符串。

*/

func countAndSay(n int) string {
	var pre = "1"
	for i := 2; i < n+1; i++ { // 从n = 2开始
		tmp := pre[0]                   // 纪录 字符串第一个字符
		count := 1                      //   字符出现了一次
		tmpResult := ""                 //
		for j := 1; j < len(pre); j++ { // 从字符串第二个字符开始遍历
			if pre[j] == tmp { // 如果当前字符等于纪录的字符
				count++ // 计数 +1
			} else {
				tmpResult += fmt.Sprintf("%d%s", count, string(tmp)) //  不等字符出现 则把计数值与字符合并
				tmp = pre[j]                                         // 纪录到新的字符
				count = 1                                            // 重置计数
			}
		}
		tmpResult += fmt.Sprintf("%d%s", count, string(tmp)) // 最后一次没有加上， 现在加上去
		pre = tmpResult
	}
	return pre
}
