package problem06

/*
题：
从尾到头打印链表算法实现

方法一：首先想到把链表的指针反转过来，在从头到尾输出
方法二：如果不允许修改链表呢，接下来想到要解决这个问题肯定要遍历链表，遍历的顺序是从头到尾，可输出顺序却要从尾到头。
也就是说，第一个遍历到的节点最后一个输出，而最后一个遍历到的节点第一个输出。这是典型的“后进先出”。我们可以用栈实现这种顺序。
每经过一个节点的时候，把该节点放到一个栈中。当遍历完整个链表后，再从栈顶开始逐个输出节点的值，此时节点的顺序已经反转过来了。
方法三：既然想到用栈来实现，而递归本质上也就是一个栈结构，于是很自然想到用递归来实现。
要实现反过来输出链表，我们每访问到一个节点，先递归输出后面的节点，在输出该节点自身，这样链表输出结果就反过来了。
用递归实现虽然看起来简洁，但是有一个问题，当链表长度非常长的时候，就会导致函数调用的层级很深，从而可能导致函数调用栈溢出。
*/
import (
	"fmt"
	"github.com/preminem/goAlgorithm/kit"
)

//  type ListNode struct {
//      Val int
//      Next *ListNode
//  }
type ListNode = kit.ListNode

// 方法二：利用栈结构迭代实现
func PrintListReversinglyIteratively(l *ListNode) {
	nodes := kit.NewStack() // 申请一个栈
	currentNode := l
	for currentNode != nil {
		nodes.Push(currentNode.Val) // 迭代链表，节点依次入栈
		currentNode = currentNode.Next
	}

	for nodes.Len() != 0 {
		value := nodes.Pop()
		fmt.Printf("%v->", value) // 出栈。打印节点值
	}
}

// 方法三：利用递归实现
func PrintListReversinglyRecurisively(l *ListNode) {
	currentNode := l
	if currentNode != nil {
		if currentNode.Next != nil {
			PrintListReversinglyRecurisively(currentNode.Next)
		}
		fmt.Printf("%v->", currentNode.Val)
	}
}
