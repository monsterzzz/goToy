package easy

func climbStairs(n int) int {
	initQue := []int{1, 2}
	for i := 2; i < n; i++ {
		initQue = append(initQue, initQue[n-1]+initQue[n-2])
	}
	return initQue[len(initQue)-1]
}
