package problem16

/*
题：
数值的整数次方
实现函数Pow(base float64, exponent int) float64，求base的exponent次方不使用库函数，同时不需要考虑大数问题
*/

func Pow(base float64, exponent int) float64 {
	if base == 0 {
		return 0
	}
	//记录指数为负的情况
	powFlag := false
	if exponent < 0 {
		powFlag = true
		exponent = -exponent
	}
	result := calPow(base, exponent)

	//指数为负时先求绝对值再求倒数
	if powFlag {
		result = 1 / result
	}
	return result
}

//a^n = a^(n/2) * a^(n/2)                  n为偶数
//      a^((n-1)/2)  *  a^((n-1)/2) * n    n为奇数
func calPow(base float64, exponent int) float64 {
	if exponent == 0 {
		return 1
	}
	if exponent == 1 {
		return base
	}
	//用右移运算符代替除以2
	result := calPow(base, exponent>>1)
	result *= result

	//用位与运算代替了求余运算符%来判断一个数是奇数还是偶数
	if exponent&1 == 1 {
		result *= base
	}

	return result
}
