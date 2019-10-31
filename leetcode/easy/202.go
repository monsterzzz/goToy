package easy

/*
编写一个算法来判断一个数是不是“快乐数”。

一个“快乐数”定义为：对于一个正整数，
每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，
也可能是无限循环但始终变不到 1。
如果可以变为 1，那么这个数就是快乐数。

*/

func isHappy(n int) bool {
	fast := n
	slow := n

	for {
		fast = squareSum(fast) // 快指针
		fast = squareSum(fast) // 慢指针
		slow = squareSum(slow)
		if fast == slow { // 如果快指针 == 慢指针  快指针==慢指针的时候说明是一个循环周期
			return fast == 1
		}
	}
}

func squareSum(n int) int {
	sum := 0
	for n > 0 {
		bit := n % 10
		sum += bit * bit
		n = n / 10
	}
	return sum
}
