package easy

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

*/

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}
	for k, v := range nums { //[3,3]
		if _, ok := m[target-v]; ok && m[target-v] != k { // map[3] = 1
			return []int{k, m[target-v]} // 0,3 -> map[6-3] = map[3]:1 -> return {0,1}
		}
	}
	return []int{}
}
