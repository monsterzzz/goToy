package easy

import "fmt"

/*
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。
*/

func Subsets(nums []int) [][]int {
	res := [][]int{{}}
	for _, v := range nums {
		//fmt.Println(v,res)
		var t [][]int
		l := len(res)

		for i := 0; i < l; i++ {

			//tmp := append(res[i], v)
			tmp := make([]int, len(res[i])+1)
			for i, vn := range res[i] {
				tmp[i] = vn
			}
			tmp[len(tmp)-1] = v

			t = append(t, tmp)
		}

		res = append(res, t...)
		fmt.Println(v, res)
		//res = append(res,[]int{})
		//res = append(res,[]int{v})
		//fmt.Println(res)
	}
	return res
}
