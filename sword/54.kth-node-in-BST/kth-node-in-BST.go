package problem54

/*
题：
二叉搜索树的第k大节点
给定一颗二叉搜索树，请找出其中第k大的节点。

分析：
二叉搜索树性质：
若它的左子树不空，则左子树上所有结点的值均小于它的根结点的值；
若它的右子树不空，则右子树上所有结点的值均大于它的根结点的值；
它的左、右子树也分别为二叉排序树。

中序遍历法很容易找出第k大节点。
 */
type BinaryTreeNode struct {
	value       int
	left, right *BinaryTreeNode
}

func kthNode(root *BinaryTreeNode, k int) *BinaryTreeNode {
	if nil == root || k <= 0 {
		return nil
	}

	var target *BinaryTreeNode
	if root.left != nil {
		target = kthNode(root.left, k)
	}

	if target == nil {
		if 1 == k {
			target = root
		} else {
			k--
		}
	}

	if target == nil && root.right != nil {
		target = kthNode(root.right, k)
	}
	return target
}
