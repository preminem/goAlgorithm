package problem28

/*
题：
对称二叉树
请实现一个函数，用来判断一颗二叉树是不是对称的。如果一个二叉树和它的镜像一样，那么他是对称的。

分析：
二叉树有三种遍历算法，前序，中序，后序，这三种都是先左节点再右节点。现在我们定义一种先右节点在左节点的遍历方法。
前序遍历 {7, 7, 7, null, null, 7, null, null, 7, 7, null, null, null} 采用我们自己定义的前序遍历{7, 7, null, 7, null. null, 7, 7, null, null, 7, null, null}
如果两个序列一样，二叉树就是对称的
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSymmetrical(root *TreeNode) bool {
	return isSymmetrical(root, root)
}

func isSymmetrical(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	if root1.Val != root2.Val {
		return false
	}

	return isSymmetrical(root1.Left, root2.Right) && isSymmetrical(root1.Right, root2.Left)
}

