package easy

import "fmt"

/*

统计所有小于非负整数 n 的质数的数量。

示例:

输入: 10
输出: 4
解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
*/

func CountPrimes(n int) int {
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	isPrime[0] = false
	isPrime[1] = false
	fmt.Println(isPrime)
	for i := 2; i < n; i++ {
		if isPrime[i] {
			for j := i * 2; j < n; j += i {
				isPrime[j] = false
			}
		}
		fmt.Println(i, isPrime)
	}

	count := 0
	for _, v := range isPrime {
		if v {
			count++
		}
	}
	return count

}

// 超时
//func CountPrimes(n int) int {
//	cur := 2
//	count := 0
//	for cur < n {
//		if isPrime(cur){
//			count++
//		}
//		cur++
//	}
//	return count
//}
//
//func isPrime(x int) bool{
//	for i:=2;i<x;i++{
//		if x % i == 0{
//			return false
//		}
//	}
//	return true
//}
