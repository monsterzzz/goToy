package easy

import "strconv"

/*
写一个程序，输出从 1 到 n 数字的字符串表示。

1. 如果 n 是3的倍数，输出“Fizz”；

2. 如果 n 是5的倍数，输出“Buzz”；

3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。

*/

func fizzBuzz1(n int) []string {
	const Fizz = "Fizz"
	const Buzz = "Buzz"
	const FizzBuzz = "FizzBuzz"
	res := make([]string, 0)
	for i := 0; i < n; i++ {
		if i < 3 {
			res = append(res, strconv.Itoa(i))
			continue
		}
		if i%3 == 0 && i%5 == 0 {
			res = append(res, FizzBuzz)
		} else if i%3 == 0 {
			res = append(res, Fizz)
		} else if i%5 == 0 {
			res = append(res, Buzz)
		} else {
			res = append(res, strconv.Itoa(i))
		}

	}
	return res
}

func fizzBuzz(n int) []string {
	m := make(map[int]string)
	m[3] = "Fizz"
	m[5] = "Buzz"
	m[15] = "FizzBuzz"
	res := make([]string, 0)
	for i := 1; i < n+1; i++ {
		res = append(res, strconv.Itoa(i))
	}
	for k, v := range m {
		for i := k - 1; i < n; i += k {
			res[i] = v
		}
	}
	return res
}
