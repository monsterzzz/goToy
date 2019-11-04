package easy

/*

给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
*/

func Reverse(x int) int {
	res := 0
	flag := false
	MAX := 2147483647
	MIN := -2147483648
	if x < 0 {
		x = x * -1
		flag = true
	}

	for {
		res *= 10
		leave := x % 10
		res += leave
		if !(MIN <= res && res <= MAX) {
			return 0
		}
		if leave == x {
			break
		}
		x = x / 10
	}

	if flag {
		return -res
	}
	return res
}
