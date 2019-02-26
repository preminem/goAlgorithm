package problem26

/*
题：
树的子结构
输入两个二叉树A和B，判断A是不是B的子结构。

分析：
1.在树A中找到和树B的根节点的值一样的节点R。
2.判断树A中以R为根节点的子树是不是包含和B一样的结构。
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func HasSubtree(rootNode1, rootNode2 *TreeNode) bool {
	result := false

	if rootNode1 != nil && rootNode2 != nil {

		// 根节点相同，判断下子树是否一致
		if rootNode1.Val == rootNode2.Val {
			result = doesTree1HaveTree2(rootNode1, rootNode2)
		}

		// 根节点不同，看子树中是否还有相同的根节点
		if !result {
			result = HasSubtree(rootNode1.Left, rootNode2)
		}

		if !result {
			result = HasSubtree(rootNode1.Right, rootNode2)
		}
	}

	return result
}


// 判断子树结构是否一致
func doesTree1HaveTree2(rootNode1, rootNode2 *TreeNode) bool {
	if rootNode2 == nil {
		return true
	}

	if rootNode1 == nil {
		return false
	}

	if rootNode1.Val != rootNode2.Val {
		return false
	}

	return doesTree1HaveTree2(rootNode1.Left, rootNode2.Left) && doesTree1HaveTree2(rootNode1.Right, rootNode2.Right)
}