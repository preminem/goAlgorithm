package problem21

/*
题：
调整数组顺序使奇数位于偶数前面
输入一个整数数组，实现一个函数来调整数组中的数字的顺序，使得所有奇数位于偶数的前半部分，所有偶数位于数组的后半部分。

分析：
[1, 2, 6, 3, 4, 7]
思路分析： 可以定义两个指针，一个在数组头，一个在数组尾。遇到奇数在偶数前面的情况交换下两数的位置即可。
对此题的程序扩展性提出要求。变得条件是奇数位于偶数前，可变为，正数位于负数前等等...。
因此考虑把这一段判断条件抽离出来封装成一个函数，后续要改变条件在定义个新的条件判断函数即可。
 */

// 把判断指针移动的的条件单独抽取出来，增加程序的可扩展性
func RecorderOddEven(array []int) []int {
	return reorder(array, isEven)
}

type fn func(int) bool

func reorder(array []int, judgeFunc fn) []int {
	length := len(array)
	if length == 0 {
		return array
	}

	pLeft := 0           // 指向数组头
	pRight := length - 1 //  指向数组尾

	for pLeft < pRight {
		// 两个指针相遇后退出

		// 指向头的指针一直向右移，直到遇到第一个偶数
		for pLeft < pRight && !judgeFunc(array[pLeft]) {
			pLeft++
		}

		// 指向尾的指针一直向左移，直到遇到第一个奇数
		for pRight > pLeft && judgeFunc(array[pRight]) {
			pRight--
		}

		// 把两个指针遇到到的数值交换
		if pLeft < pRight {
			array[pLeft], array[pRight] = array[pRight], array[pLeft]
		}
	}
	return array
}

// 判断是否为奇数
func isEven(n int) bool {
	return (n & 0x1) == 0
}
