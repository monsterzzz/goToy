package easy

/*
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
输入: [2,2,1]
输出: 1
输入: [4,1,2,1,2]
输出: 4
*/

func singleNumber(nums []int) int {
	/*
			交换律：a ^ b ^ c <=> a ^ c ^ b

		任何数于0异或为任何数 0 ^ n => n

		相同的数异或为0: n ^ n => 0

		var a = [2,3,2,4,4]

		2 ^ 3 ^ 2 ^ 4 ^ 4等价于 2 ^ 2 ^ 4 ^ 4 ^ 3 => 0 ^ 0 ^3 => 3

	*/
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res = res ^ nums[i]
	}
	return res
}
