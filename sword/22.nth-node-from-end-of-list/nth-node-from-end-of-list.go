package problem22

/*
题：
链表中倒数第k个节点
输入一个单向链表，输出该链表中倒数第k个节点。
类似于leetcode19题：删除链表的倒数第 n 个节点。

分析：
定义两个指针指向链表头部。第一个指针从头向尾移动k-1步，第二个指针保持不动。
从第k步开始，第二个指针也开始向尾部移动，由于两个指针位置保持在k-1，当第一个指针到尾部时，第二个指针移动到了第n-k+1的位置，即倒数第k个节点。

保证鲁棒性，要注意三点：
1.输入的head为空
2.节点总数小于k
3.输入的参数k<=0
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func FindKthToTail(head *ListNode, k int) (pSecond *ListNode) {
	if head == nil || k <= 0 {
		return nil
	}
	pFirst := head
	pSecond = head

	for i := 0; i < k-1; i++ {
		if pFirst.Next != nil {
			pFirst = pFirst.Next
		} else {
			return nil
		}
	}

	for pFirst.Next != nil {
		pFirst = pFirst.Next
		pSecond = pSecond.Next
	}
	return
}
