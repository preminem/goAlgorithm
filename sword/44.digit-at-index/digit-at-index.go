package problem44

/*
题：
数字序列中某一位的数字
数字以0123456789101112131415...的格式序列化到一个字符序列中。在这个序列中，第5位（从0开始计数）是5，第13位是1。请写出一个函数，求任意第n位对应数字。

分析：
以第15位数字2为例（2隶属与12，两位数，位于12从左侧以0号开始下标为1的位置）
步骤1：首先确定该数字是属于几位数的;
      如果是一位数，n<9;如果是两位数，n<9+90*2=189;
      说明是两位数。
步骤2：确定该数字属于哪个数。10+(15-10)/2= 12。
步骤3：确定是该数中哪一位。15-10-(12-10)*2 = 1, 所以位于“12”的下标为1的位置，即数字2。

以第1001位数字7为例
步骤1：首先确定该数字是属于几位数的;
      如果是一位数，n<9;如果是两位数，n<9+90*2=189;如果是三位数，n<189+900*3=2889;
      说明是三位数。
步骤2：确定该数字属于哪个数。100+(1001-190)/3= 370。
步骤3：确定是该数中哪一位。1001-190-(370-100)*3 = 1,所以位于“370”的下标为1的位置，即数字1。
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
	//确定该数字属于哪个数
	number := beginNumber(digits) + index/digits
	//确定第几位
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
		//确定该数字是几位数
		if index < numbers*digit {
			return digitAtIndexCore(index, digit)
		}
		//不是digit位数则进入下一位
		index -= digit * numbers
		digit++
	}
	return -1
}