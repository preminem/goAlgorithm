package problem0029

/*
Topic:
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
返回被除数 dividend 除以除数 divisor 得到的商。

Example:
示例 1:
输入: dividend = 10, divisor = 3
输出: 3

示例 2:
输入: dividend = 7, divisor = -3
输出: -2

说明:
被除数和除数均为 32 位有符号整数。
除数不为 0。
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31,2^31−1]。本题中，如果除法结果溢出，则返回2^31 − 1。

Thought:
不能使用乘法，除法和取余运算，编写一个除法函数。大的步骤是
取出两个数的符号和绝对值
使用绝对值做除法
还原结果的符号
检查是否溢出
其中，使用绝对值做除法是重点，思路参考程序注释。
 */

import (
	"math"
)

func divide(m, n int) int {
	// 防止有人把0当做除数
	if n == 0 {
		return math.MaxInt32
	}

	signM, absM := analysis(m)
	signN, absN := analysis(n)

	res, _ := d(absM, absN, 1)

	// 修改res的符号
	if signM != signN {
		res = res - res - res
	}

	// 检查溢出
	if res < math.MinInt32 || res > math.MaxInt32 {
		return math.MaxInt32
	}

	return res
}

func analysis(num int) (sign, abs int) {
	sign = 1
	abs = num
	if num < 0 {
		sign = -1
		abs = num - num - num
	}

	return
}

// d 计算m/n的值，返回结果和余数
// m >= 0
// n > 0
// count == 1, 代表初始n的个数，在递归过程中，count == 当前的n/初始的n
func d(m, n, count int) (res, remainder int) {
	switch {
	case m < n:
		return 0, m
	case n <= m && m < n+n:
		return count, m - n
	default:
		res, remainder = d(m, n+n, count+count)
		if remainder >= n {
			return res + count, remainder - n
		}

		return
	}
}

// 以下为上述递归方法的普通实现方式
// func d(m, n int) int {
// 	res := 0
// 	rs, ress := []int{n}, []int{1}
// 	temp, i := n+n, 1

// 	for temp <= m {
// 		rs = append(rs, temp)
// 		ress = append(ress, ress[i-1]+ress[i-1])

// 		temp += temp
// 		i++
// 	}

// 	for i := len(rs) - 1; i >= 0; i-- {
// 		if m >= rs[i] {
// 			m -= rs[i]
// 			res += ress[i]
// 		}
// 	}

// 	return res
// }


