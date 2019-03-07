package problem46

/*
题：
把数字翻译成字符串
给定一个数字，按照如下规则把它翻译成字符串：0翻译成"a",1翻译成"b"，...25翻译成"z"。一个数字可能有多个翻译。例如12258有五种不同的翻译，
分别是"bccfi"、"bwfi"、"bczi"、"mcfi"、和"mzi"。实现一个函数，用来计算一个数字有多少种不同的翻译方法。

分析：
定义函数f(i)表示从第i位数字开始的不同翻译的数目，那么f(i)=f(i+1)+g(i,i+1)*f(i+2).
当第i位和第i+1位两位数字拼接起来的数字在10~25范围内时，函数g(i,i+1)的值为1；否则为0.

自上而下：用递归思路会存在重复的子问题。
自下而上：用循环解决，可以消除重复的子问题。我们从右到左翻译并计算不同翻译的数目。
*/
import "strconv"

func getTranslation(number int) int {
	if 0 > number {
		return 0
	}
	return getTranslationCore(strconv.Itoa(number))
}

//动态规划，从右到左计算。
func getTranslationCore(str string) int {
	length := len(str)
	//count[num]表示从第num位数字开始的不同翻译的数目
	counts := make([]int, length)
	num := 0

	for i := length - 1; i >= 0; i-- {
		num = 0
		if i < length-1 {
			num = counts[i+1]
		} else {
			num = 1
		}

		if i < length-1 {
			//golang里单引号只能有一个字符，如果输出会返回这个字符的ascii码
			//-'0'是为了将字符1，2，3...转换为数字1,2,3...
			digit1 := str[i] - '0'
			digit2 := str[i+1] - '0'
			number := digit1*10 + digit2
			if 10 <= number && 25 >= number {
				if i < length-2 {
					num += counts[i+2]
				} else {
					num += 1
				}
			}
		}
		counts[i] = num
	}
	return counts[0]
}
