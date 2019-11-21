package easy

import "fmt"

/*
给定一组字符，使用原地算法将其压缩。

压缩后的长度必须始终小于或等于原数组长度。

数组的每个元素应该是长度为1 的字符（不是 int 整数类型）。

在完成原地修改输入数组后，返回数组的新长度。

*/

func compress(chars []byte) int {
	p1 := 0
	p2 := 0

	for p1 < len(chars) && p2 < len(chars) {
		if chars[p1] == chars[p2] {
			p2++
		} else {
			diff := p2 - p1
			fmt.Println(diff)
			fmt.Println(chars[p2:])
			if diff > 0 {
				tmp := make([]int, 0)
				for diff > 0 {
					tmp = append([]int{diff % 10}, tmp...)
					diff = diff / 10
				}

				for i, v := range tmp {
					chars[p1+i+1] = byte(v + 48)
				}
				chars = append(chars[:p1+len(tmp)+1], chars[p2:]...)
				p1 = p2

			}
			p1++
		}

	}
	diff := p2 - p1
	fmt.Println(diff)
	if diff > 1 {
		tmp := make([]int, 0)
		for diff > 0 {
			tmp = append([]int{diff % 10}, tmp...)
			diff = diff / 10
		}

		for i, v := range tmp {
			chars[p1+i+1] = byte(v + 48)
		}
		chars = append(chars[:p1+len(tmp)+1], chars[p2:]...)
	}
	//TODO
	fmt.Println(string(chars))
	return len(chars)
}
