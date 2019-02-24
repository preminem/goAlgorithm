package problem17

import "fmt"

/*
题：
打印从1到最大的n位数
输入数字n,按顺序打印从1到最大的n位十进制数，比如输入3，则打印输出1，2，3一直到最大的3位数999.

分析：
主要要注意n位数，n可以很大如果直接用n去循环打印，可能会溢出。

如果在数字面前补0，会发现n位所有十进制数其实就是n个从0~9的全排列。我们把数字的每一位从0到9排列一遍，就得到了所有十进制数。
只是在打印的时候，排在前面的0不打印出来。
*/

func Print1ToMaxOfDigits(n int) {
	if n <= 0 {
		return
	}
	//用来存储需要打印的数
	number := make([]int, n)

	//第一位的所有情况
	for i := 0; i < 10; i++ {
		number[0] = i
		Print1ToMaxOfNDigitsRecursively(number, 0)
	}

}

//数字的每一位都可能是0~9中的一个数，然后设置下一位
func Print1ToMaxOfNDigitsRecursively(number []int, index int) {
	length := len(number)
	//递归的结束条件是已经设置了数字的最后一位
	if index == length-1 {
		printNumber(number)
		return
	}
	//设置下一位
	for i := 0; i < 10; i++ {
		number[index+1] = i
		Print1ToMaxOfNDigitsRecursively(number, index+1)
	}
}

func printNumber(number []int) {
	isBeginning := true
	length := len(number)

	for i := 0; i < length; i++ {
		if isBeginning && number[i] != 0 {
			isBeginning = false
		}
		if !isBeginning {
			fmt.Printf("%d", number[i])
		}
	}
	fmt.Printf("\n")
}
