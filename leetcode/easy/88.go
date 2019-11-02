package easy

/*
给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。

说明:

初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。


*/

// TODO
func merge(nums1 []int, m int, nums2 []int, n int) {

	p1 := 0
	p2 := 0
	for p2 < len(nums2) {
		if nums1[p1] > nums2[p2] {
			nums1 = append([]int{nums2[p2]}, nums1...)
			p2++
		}
		p1++
	}

}
