package problem65

/*
题目：
写一个函数，求两个整数之和，要求在函数体内不得使用+、-、*、/四则运算符号。

分析：
二进制位运算分步进行：
两个数相加不进位，则相当于进行位异或运算^（0+1=1,1+0=1,0+0=0,1+1=0）；
两个数只看进位，则相当于进行位与运算&（0&0=0、1&1=1、0&1=0、1&0=0），再将进位前进一位；
前两步结果相加。
*/

func Add(num1, num2 int) int {
	var sum, carry int
	for {
		sum = num1 ^ num2
		carry = (num1 & num2) << 1
		num1 = sum
		num2 = carry

		if 0 == num2 {
			break
		}
	}
	return num1
}
