package easy

/*
给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在众数。
*/

func majorityElement1(nums []int) int {
	m := make(map[int]int) // 哈希表
	for _, v := range nums {
		if m[v]+1 > len(nums)/2 {
			return v
		}
		m[v]++
	}
	return 0
}

func majorityElement(nums []int) int { // 投票
	count := 1
	cur := nums[0]
	for _, v := range nums[1:] {
		if count == 0 {
			cur = v
			count = 1
		} else if cur == v {
			count++
		} else {
			count--
		}
	}
	return cur
}
