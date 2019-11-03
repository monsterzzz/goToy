package easy

import "fmt"

/*

给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: [1,2,3,4,5,6,7] 和 k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]
示例 2:

输入: [-1,-100,3,99] 和 k = 2
输出: [3,99,-1,-100]
解释:
向右旋转 1 步: [99,-1,-100,3]
向右旋转 2 步: [3,99,-1,-100]
说明:

尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
要求使用空间复杂度为 O(1) 的 原地 算法。


*/

func Rotate1(nums []int, k int) {
	for k != 0 {
		last := nums[len(nums)-1]
		pre := nums[0]
		var tmp int
		fmt.Println(nums)
		for j := 1; j < len(nums); j++ {

			tmp = nums[j]
			nums[j] = pre
			pre = tmp
		}
		nums[0] = last
		k--
	}
}

func Rotate(nums []int, k int) {
	k %= len(nums)
	fmt.Println(k)
	reverse(nums)
	reverse(nums[k:])
	reverse(nums[:k])
}

func reverse(nums []int) {
	p1 := 0
	p2 := len(nums) - 1

	for p1 != len(nums)/2 {
		nums[p1], nums[p2] = nums[p2], nums[p1]
		p2--
		p1++
	}
}
