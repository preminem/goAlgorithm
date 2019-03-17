package problem68

import "container/list"

/*
题目：
输入两个树节点，求他们的最低公共祖先。

二叉树（二叉搜索树）：二叉搜索树是排序的，如果当前节点大于两个节点的值，去左子树中寻找；
如果当前节点小于两个节点的值，去右子树中寻找；如果当前节点位于两个节点值之间，则该节点就是要寻找的最低公共祖先。

普通树（存在指向父节点的指针）：从给定节点出发，由父节点指针回到到根结点，形成链表。从而将问题转化为求两个链表的第一个公共节点的问题。

普通树（不存在指向父节点的指针）：利用两个辅助链表通过递归遍历的方法找到两条到达给定节点的路径，寻找两个链表最后一个公共节点，就是最低公共祖先。
*/

// 如果树是二叉搜索树
type TreeNodeA struct {
	Val   int
	Left  *TreeNodeA
	Right *TreeNodeA
}

func getLastCommentNodeA(node1, node2, root *TreeNodeA) *TreeNodeA {
	if nil == root || nil == node1 || nil == node2 {
		return nil
	}
	if node1.Val < root.Val && node2.Val < root.Val {
		return getLastCommentNodeA(node1, node2, root.Left)
	}
	if node1.Val > root.Val && node2.Val > root.Val {
		return getLastCommentNodeA(node1, node2, root.Right)
	}
	return root
}

// 带有父指针的普通树
type TreeNodeB struct {
	Val    int
	Child  []*TreeNodeB
	Parent *TreeNodeB
}

func getLastCommentNodeB(node1, node2 *TreeNodeB) *TreeNodeB {
	len1 := getLengthToRoot(node1)
	len2 := getLengthToRoot(node2)
	if len1 > len2 {
		for i := 0; i < len1-len2; i++ {
			node1 = node1.Parent
		}
	} else if len1 < len2 {
		for i := 0; i < len2-len1; i++ {
			node2 = node2.Parent
		}
	}
	for node1 != nil && node2 != nil {
		if node1.Parent == node2.Parent {
			return node1.Parent
		}
		node1 = node1.Parent
		node2 = node2.Parent
	}
	return nil
}

func getLengthToRoot(node *TreeNodeB) int {
	if node == nil {
		return 0
	}
	return 1 + getLengthToRoot(node.Parent)
}

//普通树 非递归实现
type TreeNodeC struct {
	Val   int
	Child []*TreeNodeC
}

func getLastComentNodeC(node1, node2, root *TreeNodeC) *TreeNodeC {
	//用两个链表分别保存从根节点到输入的两个节点的路径
	path1 := list.New()
	path2 := list.New()
	if nil == root || nil == node1 || nil == node2 {
		return nil
	}
	getPath(node1, root, path1)
	getPath(node2, root, path2)
	var lastParentNode *TreeNodeC
	for path1.Front() != nil && path2.Front() != nil {
		if path1.Front().Value == path2.Front().Value {
			lastParentNode = path1.Front().Value.(*TreeNodeC)
		}
		path1.Remove(path1.Front())
		path2.Remove(path2.Front())
	}
	return lastParentNode
}

//用前序遍历法得到路径
func getPath(node, root *TreeNodeC, path *list.List) bool {
	if node == root {
		return true
	}
	path.PushBack(root)
	found := false
	for i := 0; i < len(root.Child) && !found; i++ {
		found = getPath(node, root.Child[i], path)
	}
	if !found {
		path.Remove(path.Back())
	}
	return found
}
