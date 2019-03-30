package problem59

/*
题目一：
给定一个数组A[]，有一个大小为w的滑动窗口，该滑动窗口从最左边滑到最后边。在该窗口中你只能看到w个数字，每次只能移动一个位置。
我们的目的是找到每个窗口w个数字中的最大值，并将这些最大值存储在数组B中。例如数组A=[1 3 -1 -3 5 3 6 ], 窗口大小w为3。

思路：
滑动窗口可以看作是一个队列；
如果每次都遍历队列去观测最大值，那么和枚举遍历的复杂度无异。所以在队列中我们保证队列的头部是窗口中的最大值，使用队列保存数组的下标；
使用窗口最大值数组保存结果；
除了前三次入队的操作外，入队操作前都要先保存当前窗口最大值，也就是队列的第一个值；
注意：如果后入队的index数字小于队列中的index数字，则后入队的依然有成为后面窗口的最大值的可能；
但是如果后入队的index数字大于队列中的index数字，则一定是后入队的数字是窗口的最大值，所以要将前面的数字移除队列；
注意最后一个窗口的最大值。
*/
import (
	"github.com/preminem/goAlgorithm/kit"
)

type Deque = kit.Deque

func maxInWindows(num []int, size int) []int {
	maxInWindow := kit.NewDeque()

	if len(num) >= size && size >= 1 {
		index := kit.NewDeque()
		for i := 0; i < size; i++ {
			//前三次入队操作
			for !index.IsEmpty() && num[i] >= num[index.Last()] {
				index.Pop()
			}
			index.Push(i)
		}
		for i := size; i < len(num); i++ {
			//先保存当前窗口的最大值，也就是队列的第一个值
			maxInWindow.Push(num[index.First()])
			//后入队的index数字大于队列中的index数字，则一定是后入队的数字是窗口的最大值，所以要将前面的数字移除队列
			for !index.IsEmpty() && num[i] >= num[index.Last()] {
				index.Pop()
			}
			//对首元素出队
			if !index.IsEmpty() && index.First() <= i-size {
				index.Shift()
			}
			//新元素入队
			index.Push(i)
		}
		//最后一个窗口的最大值
		maxInWindow.Push(num[index.First()])
	}
	var ans []int
	for !maxInWindow.IsEmpty() {
		ans = append(ans, maxInWindow.Shift())
	}
	return ans
}
