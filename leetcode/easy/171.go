package easy

/*
给定一个Excel表格中的列名称，返回其相应的列序号。

例如，

    A -> 1
    B -> 2
    C -> 3
    ...
    Z -> 26
    AA -> 27
    AB -> 28
    ...
*/

func titleToNumber(s string) int {
	res := 0
	for _, v := range s {
		num := int(v) - int('A') + 1 // Z - A = 25  25 + 1 = 26
		// Y - A = 24 24 + 1 = 25
		res = res*26 + num // res = 0 * 26 +26 = 26
		// res = 26 * 26 + 25 = 701
	}

	return res
}
