package easy

/*
编写一个函数，输入是一个无符号整数，返回其二进制表达式中数字位数为 ‘1’ 的个数
*/

func hammingWeight(num uint32) int {
	s := 0
	for num != 0 {
		s++
		num &= num - 1
	}
	return s
}
func HammingWeight(num uint32) int {
	s := 0
	for num != 0 {
		s++
		num &= num - 1
	}
	return s
}
