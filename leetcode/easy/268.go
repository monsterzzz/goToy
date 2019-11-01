package easy

/*
给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数。

示例 1:

输入: [3,0,1]
输出: 2
示例 2:

输入: [9,6,4,2,3,5,7,0,1]
输出: 8
说明:
你的算法应具有线性时间复杂度。你能否仅使用额外常数空间来实现?

*/

func missingNumber(nums []int) int {

	res := len(nums)

	for i := 0; i < len(nums); i++ { // len(nums) = 9  最大值为9
		res ^= i ^ nums[i]
	}

	// 0 ^ 0 ^ 1 ^ 1 ^ 2 ^ 2 ^ 3 ^ 3 ^ 4 ^ 4 ^ 5 ^ 5 ^ 6 ^ 6 ^ 7 ^ 7 ^ 8  ^ 9 ^ 9
	// 0 ^ 0 ^ 0 ^ 0 ^ 0 ^ 0 ^ 0 ^ 0 ^ 8 ^ 0
	// 8

	// 0 ^ 0 ^ 1 ^ 1 ^ 2 ^ 2 ^ 3 ^ 3 ^ 4 ^ 4 ^ 5 ^ 5 ^ 6 ^ 6 ^ 7 ^ 7 ^ 8  ^ 9 ^ 9
	return res

}