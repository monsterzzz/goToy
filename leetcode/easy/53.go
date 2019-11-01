package easy

func maxSubArray(nums []int) int {
	var ans = nums[0]
	var sum = 0

	// 如果当前总和是负数那么如果下一个值比你当前总和大那干脆就从下一个地方重新开始，否则就可以一直加下去。

	//TODO

	for _, v := range nums {
		if sum <= 0 {
			sum = v
		} else {
			sum += v
		}

		if ans < sum {
			ans = sum
		}
	}
	return ans
}
