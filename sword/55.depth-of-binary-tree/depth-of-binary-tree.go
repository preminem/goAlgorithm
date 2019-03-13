package problem55

/*
题目一：二叉树的深度
输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。

分析：
使用递归方法，树的深度是其左、右子树深度较大值+1.
 */

type BinaryTreeNode struct {
	value       int
	left, right *BinaryTreeNode
}

func treeDepth(root *BinaryTreeNode) int {
	if nil == root {
		return 0
	}
	left := treeDepth(root.left)
	right := treeDepth(root.right)
	if left >= right {
		return left + 1
	} else {
		return right + 1
	}
}

/*
题目二：平衡二叉树
输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。

分析：
1.这个题目需要清除的指导平衡二叉树的定义：它是一棵空树或它的左右两个子树的高度差的绝对值不超过1，并且左右两个子树都是一棵平衡二叉树。
2.使用后序遍历，先判断左右子树是不是平衡的，如果是平衡的那么就分别记录左右子树的高度并返回，然后判断当前根节点是否是平衡的。
 */

func isBalanced(root *BinaryTreeNode, depth *int) bool {
	if nil == root {
		*depth = 0
		return true
	}
	var left, right int
	if isBalanced(root.left, &left) && isBalanced(root.right, &right) {
		diff := left - right
		if -1 <= diff && diff <= 1 {
			if left > right {
				*depth = 1 + left
			} else {
				*depth = 1 + right
			}
			return true
		}
	}
	return false
}

func isBalance(root *BinaryTreeNode) bool {
	var depth int
	return isBalanced(root, &depth)
}