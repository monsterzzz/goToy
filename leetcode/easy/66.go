package easy

func plusOne(digits []int) []int {

	/*

		如果当前位加一不等于10,那么直接返回
		如果等于10
			把当前位取0
			如果当前位还是首位 那么在首位加1

	*/
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += 1
		if digits[i] != 10 {
			break
		}
		digits[i] = 0
		if i == 0 {
			digits = append([]int{1}, digits...)
		}

	}

	return digits
}
