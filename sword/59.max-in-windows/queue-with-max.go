package problem59

/*
题目二：
队列的最大值
请定义一个队列并实现函数max得到队列里的最大值，要求函数max、push_back和pop_front的时间复杂度都是O(1)
此题的意思就是，假如有一家很火爆的店，店外不断有人排队离队，请你用队列这个数据结构展示目前排队的状态，即任何时刻队列里的最大值（可以想象成队列里最高的人）

思路：
同上一题相同，我们要寻找队列的最大值，相当与将滑动窗口设置为整个队列。
这里需要使用两个队列，一个队列用来保存入队的数据，一个队列用来保存队列的当前最大值。
其中两个队列中的数据都是一对值（元素值和其索引）
同时需要注意出队操作，数据队列出队的同时需要判断其索引是否和当前最大值队列首部索引相同，如果相同则同时也将最大值队列头部出队。
*/
import (
	"github.com/preminem/goAlgorithm/kit"
)

type Queue = kit.Queue

type MyQueue struct {
	Queue *Queue
	Deque *Deque
}

func NewMyQueue() *MyQueue {
	return &MyQueue{kit.NewQueue(), kit.NewDeque()}
}

func (q *MyQueue) Push(n int) {
	for !q.Deque.IsEmpty() && n >= q.Deque.Last() {
		q.Deque.Pop()
	}
	q.Queue.Push(n)
	q.Deque.Push(n)
}

func (q *MyQueue) Pop() int {
	if q.Deque.IsEmpty() {
		panic("queue is empty")
	}
	if q.Deque.First() == q.Queue.Peek() {
		q.Deque.Pop()
	}
	return q.Queue.Pop()
}

func (q *MyQueue) max() int {
	if q.Deque.IsEmpty() {
		panic("queue is empty")
	}
	return q.Deque.First()
}
