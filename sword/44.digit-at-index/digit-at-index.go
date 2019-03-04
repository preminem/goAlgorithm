package problem

/*
题：
数字序列中某一位的数字
数字以0123456789101112131415...的格式序列化到一个字符序列中。在这个序列中，第5位（从0开始计数）是5，第13位是1。请写出一个函数，求任意第n位对应数字。

 */
import "math"

//m位的数字总共有多少个
func countOfIntegers(digits int) int {
	if 1 == digits {
		return 10
	}
	return int(9 * math.Pow10(digits-1))
}

//第一个m位数
func beginNumber(digits int) int {
	if 1 == digits {
		return 0
	}
	return int(math.Pow10(digits - 1))
}

func digitAtIndexCore(index, digits int) int {
	number := beginNumber(digits) + index/digits
	indexFromRight := digits - index%digits
	for i := 1; i < indexFromRight; i++ {
		number /= 10
	}
	return number % 10
}

func digitAtIndex(index int) int {
	if 0 > index {
		return -1
	}
	digit := 1
	for {
		numbers := countOfIntegers(digit)
		if index < numbers*digit {
			return digitAtIndexCore(index, digit)
		}
		index -= digit * numbers
		digit++
	}
	return -1
}