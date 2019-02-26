package problem23

/*
题：
链表中环的入口节点
如果一个链表中包含环，如何找出环的入口节点

分析：
首先确定链表存在环
可以定义两个指针，都指向链表头。第一个以1的步速向尾部移，第二个以2的步速向尾部移。如果两个在A节点相遇，则说明存在环。

确定存在环后，假设环的长度为n。在定义两个指针后，第一个指向头部，第二个指向距离第一个节点n的节点。
两个节点在以相同的速度移动。当两个节点相遇时，相遇的节点就是环的入口节点

现在问题转化为求环的长度。
第1步中，两个指针相遇点A一定在环中， 继续以1的步速移动，再回到A节点时，移动的距离就是环的长度。
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 寻找环的入口节点
func EntryNodeOfLoop(head *ListNode) *ListNode {
	length := LengthOfLoop(head)

	// 不存在环
	if length == 0 {
		return nil
	}

	// 申明两个指针，之间差了环的长度个节点
	p1 := head
	p2 := p1
	for i := 0; i < length; i++ {
		p2 = p2.Next
	}

	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p1

}

// 计算环的长度
func LengthOfLoop(head *ListNode) int {
	p1 := head.Next // 移动慢的指针，一次一步
	p2 := p1.Next   // 移动快的指针，一次两步
	length := 0          // 环的长度
	hasLoop := false

	for p1 != nil && p2 != nil {
		//p2追上p1, 说明存在环
		if p1 == p2 {
			hasLoop = true
			break
		}

		p1 = p1.Next
		p2 = p2.Next
		if p2 != nil {
			p2 = p2.Next
		}
	}

	// 存在环的话，此时p2不动，p1在以1的步速移动。
	// 统计移动次数，直到与p2相遇,移动的次数即为环大小。
	if hasLoop {
		length = 1
		p1 = p1.Next
		for p1 != p2 {
			length++
			p1 = p1.Next
		}
	}
	return length
}
