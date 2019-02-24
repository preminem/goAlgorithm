package problem18

/*
题：
删除链表的节点
*/

import "github.com/preminem/goAlgorithm/kit"

type ListNode = kit.ListNode

//题目一：在O(1)时间内删除链表节点。给定单向链表的头指针和一个节点指针，定义一个函数在O(1)时间内删除该节点。
func delete(l *ListNode, pToBeDeleted *ListNode) {
	pHead := l
	//把下一个节点的内容复制到需要删除的节点上覆盖原有内容，再把下一个节点删除
	if pToBeDeleted.Next != nil {
		pToBeDeleted = pToBeDeleted.Next
		//只有一个节点，且要删除
	} else if pHead.Val == pToBeDeleted.Val {
		l = nil
		//要删除的节点位于链表尾部时只能顺序遍历
	} else {
		cur := pHead
		for cur.Next.Val != pToBeDeleted.Val {
			cur = cur.Next
		}
		cur.Next = nil
	}
}

//题目二：删除链表中的重复节点。在一个排序的链表中，1->2->3->3->4->4->5，删除重复节点后1->2->5
func deleteDuplicates(head *ListNode) *ListNode {
	// 长度 <=1 的 list ，可以直接返回
	if head == nil || head.Next == nil {
		return head
	}

	// 要么 head 重复了，那就删除 head
	if head.Val == head.Next.Val {
		for head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
		}
		return deleteDuplicates(head.Next)
	}

	// 要么 head 不重复，递归处理 head 后面的节点
	head.Next = deleteDuplicates(head.Next)
	return head
}
