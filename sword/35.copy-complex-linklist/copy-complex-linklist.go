package problem35

/*
题：
复杂链表的复制
请实现一个函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个m_pNext指针指向下一个节点，还有一个m_pSibling指针指向链表中的任意节点或者nullptr

分析：
分成三步
1.先不管pSibling，复制链表，复制的节点插在对应节点的后面
2.复制pSibling的对应关系
3.把复制的链表拆分出来
 */

func CloneNodes(head *ComplexListNode) *ComplexListNode {
	cloneNodes(head)
	connectSiblingNodes(head)
	return reconnectNodes(head)
}

// 第一步复制该节点，插入在该节点和该节点的下一个节点之间
func cloneNodes(head *ComplexListNode) {
	node := head
	for node != nil {
		clonedNode := &ComplexListNode{}
		clonedNode.value = node.value
		clonedNode.next = node.next
		clonedNode.sibling = nil

		node.next = clonedNode
		node = clonedNode.next
	}
}

// 第二步，复制链表的兄弟关系
func connectSiblingNodes(head *ComplexListNode) {
	node := head
	for node != nil {
		cloned := node.next
		if node.sibling != nil {
			cloned.sibling = node.sibling.next
		}
		node = cloned.next

	}
}

// 把链表拆成两个
func reconnectNodes(head *ComplexListNode) *ComplexListNode {
	node := head
	var clonedHead, clonedNode *ComplexListNode

	if node != nil {
		clonedHead = node.next
		clonedNode = node.next
		node.next = clonedNode.next
		node = node.next
	}

	for node != nil {
		clonedNode.next = node.next
		clonedNode = clonedNode.next

		node.next = clonedNode.next
		node = node.next
	}

	return clonedHead
}