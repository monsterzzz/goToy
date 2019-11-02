package main

import "fmt"

func main() {

	numRows := 5

	//if numRows == 1 {
	//	return [][]int{{1}}
	//}else if numRows == 2{
	//	return [][]int{{1},{1,1}}
	//}else if numRows == 0{
	//	return [][]int{}
	//}

	var res = [][]int{{1}, {1, 1}}
	for i := 2; i < numRows; i++ {
		tmp := make([]int, len(res[i-1])+1)
		tmp[0] = 1
		tmp[len(tmp)-1] = 1
		for j := 1; j < len(res[i-1]); j++ {
			cur := res[i-1][j-1] + res[i-1][j]
			tmp[j] = cur
		}
		res = append(res, tmp)
	}
	fmt.Println(res)

}
