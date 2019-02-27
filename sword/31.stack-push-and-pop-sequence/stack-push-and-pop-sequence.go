package problem31

/*
题：
输入两个整数序列，第一个表示压栈顺序，判断第二个序列是否为该栈的弹出序列？（压入栈的数字均不相等)

分析：
因为有入栈序列，所以从入栈开始，遍历入栈序列
1. 遍历入栈序列时，每遍历一个，就判断是否是出栈序列中的需要弹出的值，
2. 如果是则弹出，然后出栈序列加一，接着继续根据入栈序列添加（for循环继续）
3. 如果不是，则继续根据入栈序列添加，直到匹配出栈序列需要的数字为止
4. 如果遍历完都发现没有匹配的，则该出栈序列不是合法出栈序列
 */

import (
	"github.com/preminem/goAlgorithm/kit"
)

func isPopOrder(pushStack, popStack []int) bool {
	length := len(pushStack)
	if 0 == length || length != len(popStack) {
		return false
	}

	stack := kit.NewStack()
	nextPush, nextPop := 0, 0

	for nextPop < length {
		//栈顶元素不等于出栈序列中的需要弹出的值，就不停取入栈序列入栈
		for stack.IsEmpty() == true || stack.Peek() != popStack[nextPop] {
			if nextPush == length {
				break
			}
			stack.Push(pushStack[nextPush])
			nextPush++

		}
		//入栈序列取完，依然不等，跳出
		if  stack.Peek() != popStack[nextPop] {
			break
		}

		//出栈，比对下一个元素
		stack.Pop()
		nextPop++
	}

	if stack.IsEmpty() && nextPop == length {
		return true
	}
	return false
}
