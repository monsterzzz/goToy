package easy

/*
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

输入: 5
输出:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]

*/

func generate(numRows int) [][]int {

	if numRows == 1 {
		return [][]int{{1}}
	} else if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	} else if numRows == 0 {
		return [][]int{}
	}

	var res = [][]int{{1}, {1, 1}}
	for i := 2; i < numRows; i++ {
		tmp := make([]int, len(res[i-1])+1)
		tmp[0] = 1
		tmp[len(tmp)-1] = 1
		for j := 1; j < len(res[i-1])-1; j++ {
			cur := res[i-1][j-1] + res[i-1][j]
			tmp[j] = cur
		}
		res = append(res, tmp)
	}
	return res

}
