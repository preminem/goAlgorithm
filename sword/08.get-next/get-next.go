package problem08

/*
题：
二叉树的下一个节点
给定一颗二叉树和其中一个节点，如何找出中序遍历序列的下一个节点？
树中的节点除了有两个分别指向左右子节点的指针，还有一个指向父节点的指针

分析：
情况一：如果一个节点有右子树，那么它的下一个节点一点就是他的右子树中最左节点。
也就是说，从右子节点出发一直沿着指向左子节点的指针，我们就能找到他的下一个节点。

情况二：如果没有右子树，如果节点是它父节点的左子节点，那么它的下一个节点就是他的父节点。

情况三：如果既没有右子树，并且是它父节点的右子节点,我们沿着指向父节点的指针一直向上遍历，直到找到一个是他父节点的左子节点的节点，
如果这样的节点存在那么，这个节点的父节点就是我们要找的节点.
*/
import "fmt"

type Node struct {
	value  int
	left   *Node
	right  *Node
	parent *Node
}

func GetNext(node *Node) *Node {
	if node == nil {
		return nil
	}
	var next *Node
	if node.right != nil { // 有右子树，直接找右子树的最左节点
		next = node.right
		for next.left != nil {
			next = next.left
		}
	} else if node.parent != nil {
		cur := node
		parent := node.parent
		for parent != nil && cur == parent.right {
			cur = parent
			parent = parent.parent
		}
		next = parent
	}
	return next
}
