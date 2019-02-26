package problem24

/*
题：
反转链表
给定一个单向链表， 把节点的指向关系反转 例如2->3->4->6->7反转后7->6->4->3->2
 */

type ListNode struct {
	Val  int
	Next *ListNode
}


func ReverseList(head *ListNode) *ListNode{
	var prev, next *ListNode
	cur := head

	for cur != nil {
		next = cur.Next
		// 掉转指针
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}