package problem49

/*
题：
我们把只包含因子2、3和5的数称作丑数（Ugly Number）。
求按从小到大的顺序的第1500个丑数。例如6、8都是丑数，但14不是，因为它包含因子7。习惯上我们把1当做第一个丑数。

分析：根据丑数的定义，我们可以知道丑数可以由另外一个丑数乘以2，3或者5得到。
因此我们可以创建一个数组，里面的数字是排好序的丑数，每一个丑数都是前面的丑数乘以2，3或者5得到的。
我们把得到的第一个丑数乘以2以后得到的大于M的结果记为M2。同样，我们把已有的每一个丑数乘以3和5，能得到第一个大于M的结果M3和M5。
那么M后面的那一个丑数应该是M2,M3和M5当中的最小值：Min(M2,M3,M5)。
比如将丑数数组中的数字按从小到大乘以2，直到得到第一个大于M的数为止，那么应该是2*2=4<M，3*2=6>M，所以M2=6。同理，M3=6，M5=10。
所以下一个丑数应该是6。
*/

func getUglyNumber(index int) int {
	if 0 >= index {
		return 0
	}

	uglyNumbers := make([]int, index)
	uglyNumbers[0] = 1
	nextUglyIndex := 1

	//用来记录下一个被乘丑数的位置
	var pMultiply2, pMultiply3, pMultiply5 int
	for nextUglyIndex < index {
		uglyNumbers[nextUglyIndex] = min3(uglyNumbers[pMultiply2]*2, uglyNumbers[pMultiply3]*3, uglyNumbers[pMultiply5]*5)

		for uglyNumbers[pMultiply2]*2 <= uglyNumbers[nextUglyIndex] {
			pMultiply2++
		}
		for uglyNumbers[pMultiply3]*3 <= uglyNumbers[nextUglyIndex] {
			pMultiply3++
		}
		for uglyNumbers[pMultiply5]*5 <= uglyNumbers[nextUglyIndex] {
			pMultiply5++
		}
		nextUglyIndex++
	}
	return uglyNumbers[index-1]
}

func min3(num1, num2, num3 int) int {
	number := min(num1, num2)
	return min(num3, number)
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
