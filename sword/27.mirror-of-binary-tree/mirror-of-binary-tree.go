package problem27

/*
题：
二叉树的镜像
完成一个函数，输入一颗二叉树，该函数输出他的镜像。
先看下树的镜像的概念,镜像就是原先二叉树照镜子后在镜子中的成像。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MirrorRecursively(node *TreeNode) {
	if node.Left == nil && node.Right == nil {
		return
	}

	node.Left, node.Right = node.Right, node.Left

	if node.Left != nil {
		MirrorRecursively(node.Left)
	}

	if node.Right != nil {
		MirrorRecursively(node.Right)
	}
}
