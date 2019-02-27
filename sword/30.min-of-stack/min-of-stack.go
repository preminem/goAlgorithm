package problem30

/*
题：
包含min函数的栈
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的min函数。在该栈中调用min(),push(),pop()的复杂度都是O(1).

分析：
把每次的最小元素（之前的最小元素和新压入栈的元素两者的较小值）都保存起来放到另一个辅助栈。
*/
import (
	"errors"
	"github.com/preminem/goAlgorithm/kit"
)

var (
	dataStack = kit.NewStack()
	minStack  = kit.NewStack()
)

func push(number int) {
	dataStack.Push(number)

	if 0 == minStack.Len() {
		minStack.Push(number)
	} else {
		topNumber := minStack.Peek()

		if number < topNumber {
			minStack.Push(number)
		} else {
			minStack.Push(topNumber)
		}
	}
}

func pop() {
	if 0 < dataStack.Len() && 0 < minStack.Len() {
		dataStack.Pop()
		minStack.Pop()
	}
}

func min() (int, error) {
	if 0 < dataStack.Len() && 0 < minStack.Len() {
		topNumber := minStack.Peek()
		return topNumber, nil
	}
	return 0, errors.New("栈为空")
}


