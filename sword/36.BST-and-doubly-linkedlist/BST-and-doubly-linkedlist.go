package problem36

/*
题：
二叉搜索树与双向链表
输入一颗二叉搜索树，将该二叉搜索树转换成一个排序的双向链表。要求不能创建任何新的节点，只能调整树中节点指针的指向。

分析：
中序遍历
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Convert(rootNode *TreeNode) *TreeNode {
	//用来指向双向链表当前最后一个节点的地址
	var lastNodeInList *TreeNode

	convertNode(rootNode, &lastNodeInList)
	headOfList := lastNodeInList
	//此时headOfList指向双向链表末端，因此要找到双向链表头部将其返回
	for headOfList != nil && headOfList.Left != nil {
		headOfList = headOfList.Left
	}

	return headOfList
}

func convertNode(node *TreeNode, lastNodeInList **TreeNode) {
	current := node

	if current.Left != nil {
		convertNode(current.Left, lastNodeInList)
	}

	// 当前节点加一个指向上一个节点的指针
	current.Left = *lastNodeInList

	if *lastNodeInList != nil {
		(*lastNodeInList).Right = current
	}

	*lastNodeInList = current
	if current.Right != nil {
		convertNode(current.Right, lastNodeInList)
	}
}
