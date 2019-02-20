package problem0024

/*
Topic:
两两交换链表中的节点
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

Example:
给定 1->2->3->4, 你应该返回 2->1->4->3.

Thought:
利用递归，简洁明了
*/
import (
	"github.com/preminem/goAlgorithm/kit"
)

// ListNode is definition for singly-linked list
type ListNode = kit.ListNode

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head

	return newHead
}