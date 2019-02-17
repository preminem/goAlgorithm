package problem0019

/*
Topic:
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

Example:
Given linked list: 1->2->3->4->5, and n = 2.
After removing the second node from the end, the linked list becomes 1->2->3->5.

Thought:
我们可以使用两个指针而不是一个指针。第一个指针从列表的开头向前移动n+1步，而第二个指针将从列表的开头出发。
现在，这两个指针被n个结点分开。我们通过同时移动两个指针向前来保持这个恒定的间隔，直到第一个指针到达最后一个结点。
此时第二个指针将指向从最后一个结点数起的第n个结点。我们重新链接第二个指针所引用的结点的next指针指向该结点的下下个结点。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	d, headIsNthFromEnd := getDaddy(head, n)

	if headIsNthFromEnd {
		// 删除head节点
		return head.Next
	}

	d.Next = d.Next.Next

	return head
}

// 获取倒数第N个节点的父节点
func getDaddy(head *ListNode, n int) (daddy *ListNode, headIsNthFromEnd bool) {
	daddy = head

	for head != nil {
		if n < 0 {
			daddy = daddy.Next
		}

		n--
		head = head.Next
	}

	// n==0，说明链的长度等于n
	headIsNthFromEnd = n == 0

	return
}
