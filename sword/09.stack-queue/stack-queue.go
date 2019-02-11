package problem09

//题：两个栈实现队列
import (
	"github.com/preminem/goAlgorithm/kit"
)

type Stack = kit.Stack

type Queue struct {
	stackA *Stack
	stackB *Stack
}

func NewQueue() *Queue {
	stackA:=kit.NewStack()
	stackB:=kit.NewStack()
	return &Queue{stackA: stackA,
		          stackB: stackB,
    }
}

//插入一个元素入栈stackA
func (q *Queue) Push(n int) {
	q.stackA.Push(n)
}

//删除一个元素时，如果stackB不为空stackB出栈。如果stackB为空时stackA先倒入stackB，stackB再出栈。
func (q *Queue) Pop() int {
	if q.stackB.Len() == 0{
		for !q.stackA.IsEmpty() {
			q.stackB.Push(q.stackA.Pop())
		}
	}
	return q.stackB.Pop()
}

//相关：用两个队列实现一个栈
