package problem0002

/*
Topic:
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8

Thought：
题目的本意是，把整数换了一种表达方式后，实现其加法。
特别注意以下情况
1.当一个列表比另一个列表长时。
2.当一个列表为空时，即出现空列表。
3.求和运算最后可能出现额外的进位。

Complexity:
时间复杂度：O(max(len(l1), len(l2))
空间复杂度：O(max(len(l1), len(l2))),新列表的长度最多为max(m,n)+1。
*/

import (
	"github.com/preminem/goAlgorithm/kit"
)

// ListNode defines for singly-linked list.
//  type ListNode struct {
//      Val int
//      Next *ListNode
//  }

type ListNode = kit.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resPre := &ListNode{}
	cur := resPre
	//carry代表进位
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10

		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}

	return resPre.Next
}
