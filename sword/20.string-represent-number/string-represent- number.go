package problem20

/*
题：
请实现一个函数用来判断字符串是否表示数值(包括整数和小数)
例如，字符串"+100" "5e2" "-123" "3.1416" "-1E-16"都是表示数值，但是"12e" "1a3.14" "1.2.3" "+-5" "12E+5.4"都不是。

分析：
表示数值的字符串遵循模式A[.[B]][e|EC]或者.B[e|EC],其中A为数值的整数部分，B在小数点后为小数部分，C在e或E后为指数部分。
在小数里可能没有数值的整数部分，例如.123等于0.123.因此A部分不是必须的。如果一个数没有整数部分，那么小数部分不能为空。
A与C都是可能以"+" 或 "-"开头的09数字字符串， B也是09的数字字符串，但前面不能有+-号
判断一个字符串是否符合上面所诉的模式时，首先尽可能多的扫描0~9的数位(有可能在起始处有"+"或"-"), 也就是前面模式中表示整数的A部分。
如果遇到"."，则开始扫描表示小数部分的B部分。
如果遇到"E"或"e"，则开始扫描表示指数的C部分。
 */

// 数字格式字符串A[.[B]][e|EC]或者.B[e|EC]表示
// 整数可以有正负，B是一个无符号整数
func IsNumeric(str string) bool {
	strLen := len(str)
	if strLen == 0 {
		return false
	}
	pIndex := 0

	// 检测一个整数部分A
	numeric := scanInteger(str, &pIndex)

	// 此时已经到字符串尾部了，把前面的结果返回
	if pIndex == strLen {
		return numeric
	}

	// 检测小数部分B
	if string(str[pIndex]) == "." {
		pIndex++

		// 下面一行代码用||的原因：
		// 1. 小数可以没有整数部分，例如.123等于0.123；
		// 2. 小数点后面可以没有数字，例如233.等于233.0；
		// 3. 当然小数点前面和后面可以有数字，例如233.666
		numeric = scanUnginnedInteger(str, &pIndex) || numeric
	}

	// 此时已经到字符串尾部了，把前面的结果返回
	if pIndex == strLen {
		return numeric
	}

	// 检测指数部分C
	if string(str[pIndex]) == "e" || string(str[pIndex]) == "E" {
		pIndex++
		if pIndex == strLen {
			return false
		}
		// 下面一行代码用&&的原因：
		// 1. 当e或E前面没有数字时，整个字符串不能表示数字，例如.e1、e1；
		// 2. 当e或E后面没有整数时，整个字符串不能表示数字，例如12e、12e+5.4
		numeric = numeric && scanInteger(str, &pIndex)

	}

	return numeric && pIndex == len(str) // 三个部分都满足，且整个字符串都匹配完了
}

func scanInteger(str string, pIndex *int) bool {
	if string(str[*pIndex]) == "+" || string(str[*pIndex]) == "-" {
		*pIndex++
	}
	return scanUnginnedInteger(str, pIndex)
}

func scanUnginnedInteger(str string, pIndex *int) bool {
	beforeIndex := *pIndex
	for i := *pIndex; i <= len(str)-1; i++ {
		if string(str[i]) >= "0" && string(str[i]) <= "9" {
			*pIndex++
		} else {
			break
		}
	}
	return *pIndex > beforeIndex
}
