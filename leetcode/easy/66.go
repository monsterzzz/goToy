package easy

func plusOne(digits []int) []int {

	for i := len(digits) - 1; i > 0; i++ {
		if (digits[i]+1)%10 == 0 {
			digits[i-1] += 1
		} else {
			digits[i] += 1
		}
	}

	return digits
}
