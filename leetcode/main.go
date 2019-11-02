package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}

	p1 := 0
	p2 := 0
	m := 6
	n := 3
	nums1 = nums1[:m]
	nums2 = nums2[:n]
	fmt.Println(len(nums1))
	for p2 < len(nums2) {
		fmt.Println(p1, p2, nums1, nums2)
		cNumS1 := make([]int, len(nums1))
		copy(cNumS1, nums1)
		if nums2[p2] <= nums1[p1] {
			nums1 = append(nums1[:p1], nums2[p2])
			nums1 = append(nums1, cNumS1[p1:]...)
			p1++
			p2++
		} else {
			if nums1[p1+1] > nums2[p2] {

			}
			nums1 = append(nums1[:p1+1], nums2[p2])
			nums1 = append(nums1, cNumS1[p1+1:]...)
			p1++
			p2++
		}
		p1++
	}
	fmt.Println(nums1)

}
