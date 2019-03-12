package problem52

/*
题：
两个链表的第一个公共节点
输入两个链表，找出它们的第一个公共节点

分析：
首先遍历两个链表得到它们的长度，就能知道哪个链表比较长，以及长的链表比短的链表多几个节点。
在第二次遍历的时候，较长的链表先走若干步，接着同时在两个链表上遍历，找到的第一个相同的节点就是它们的第一个公共节点。
 */

type ListNode struct {
	value int
	next  *ListNode
}

func findFirstCommonNode(head1, head2 *ListNode) *ListNode {
	length1 := getListLength(head1)
	length2 := getListLength(head2)
	var lengthDif int
	var listHeadLong, listHeadShort, firstCommonNode *ListNode
	if length1 >= length2 {
		lengthDif = length1 - length2
		listHeadLong = head1
		listHeadShort = head2
	} else {
		lengthDif = length2 - length1
		listHeadLong = head2
		listHeadShort = head1
	}

	for i := 0; i < lengthDif; i++ {
		listHeadLong = listHeadLong.next
	}
	for listHeadLong != listHeadShort && nil != listHeadLong && nil != listHeadShort {
		listHeadLong = listHeadLong.next
		listHeadShort = listHeadShort.next
	}
	firstCommonNode = listHeadLong
	return firstCommonNode
}

func getListLength(head *ListNode) int {
	var length int
	node := head
	for nil != node {
		length++
		node = node.next
	}
	return length
}