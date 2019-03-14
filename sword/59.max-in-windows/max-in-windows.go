package problem59

/*
题目一：
给定一个数组A[]，有一个大小为w的滑动窗口，该滑动窗口从最左边滑到最后边。在该窗口中你只能看到w个数字，每次只能移动一个位置。
我们的目的是找到每个窗口w个数字中的最大值，并将这些最大值存储在数组B中。例如数组A=[1 3 -1 -3 5 3 6 ], 窗口大小w为3。

思路：
利用一个双端队列，保存滑动窗口中的有可能成为最大值的下标。然后随着滑动窗口的移动，不断跟更新队尾和队首.
而每次窗口移动都将队首压入最大值数组。
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
			for !index.IsEmpty() && num[i] >= num[index.Last()] {
				index.Pop()
			}
			index.Push(i)
		}
		for i := size; i < len(num); i++ {
			maxInWindow.Push(num[index.First()])
			for !index.IsEmpty() && num[i] >= num[index.Last()] {
				index.Pop()
			}
			if !index.IsEmpty() && index.First() <= i-size {
				index.Shift()
			}
			index.Push(i)
		}
		maxInWindow.Push(num[index.First()])
	}
	var ans []int
	for !maxInWindow.IsEmpty() {
		ans = append(ans, maxInWindow.Shift())
	}
	return ans
}
