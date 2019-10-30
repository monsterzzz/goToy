package easy

/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
*/

func MoveZeroes(nums []int) {
	pos := -1 // 纪录最开始0的位置
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 && pos == -1 { // 如果遇到第一个0
			pos = i
		} else if pos != -1 && nums[i] != 0 { // 其余非0数
			nums[i], nums[pos] = nums[pos], nums[i] // 交换位置
			pos++                                   // 0向后移动
		}
	}
}
