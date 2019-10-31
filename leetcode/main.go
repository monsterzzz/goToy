package main

import "fmt"

func main() {

	var pre = "1"
	n := 4
	for i := 2; i < n+1; i++ {
		tmp := pre[0]
		count := 1
		tmpResult := ""
		for j := 1; j < len(pre); j++ {
			if pre[j] != tmp {
				tmpResult += fmt.Sprintf("%d%s", count, string(tmp))
				tmp = pre[j]
				count = 1
			} else {
				count++
			}
		}
		tmpResult += fmt.Sprintf("%d%s", count, string(tmp))
		fmt.Println(i, pre, tmpResult)
		pre = tmpResult
	}

}
