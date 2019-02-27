package problem33

/*
题：
二叉搜索树的后序遍历序列
输入一个整数数组，判断该数组是不是某个二叉搜索树的后序遍历结果。如果是则返回true，否则返回false。
假设输入的数组任意两个数字都是不相同的 例如输入数组{5, 7, 6, 9, 11, 10, 8}则返回true，它是下面二叉搜索树的后续遍历的结果

分析：
分析:

后续遍历序列特点

最后一个数字是树的根节点值。
数组中前面的数字可以分成两部分：第一部分是左子树节点的值，它们都比根节点的值小。第二部分是右子树的值，它们都比根节点的值大。
以数组{5, 7, 6, 9, 11, 10, 8}为例子， 8是根节点的值。在这个数组中，5, 7, 6的值都比8小。是8这个节点的左子树。9，11，10都比8这个节点值大，是节点8的右子树。

然后以同样的方法确定数组中每一部分的子树结构。对于序列{5,7,6}。6为根的值。5为左子树，7为右子树。
 */

func VerifySquenceOfBST(sequence []int) bool {
	length := len(sequence)
	if length == 0 {
		return false
	}

	root := sequence[length-1]

	var i, j int
	// 在二叉搜索树中左子树的结点小于根结点
	for i = 0; i < length-1; i++ {
		if sequence[i] > root {
			break
		}
	}

	// 在二叉搜索树中右子树的结点大于根结点
	for j = i; j < length-1; j++ {
		if sequence[j] < root {
			return false
		}
	}

	left := true
	if i > 0 {
		// 判断左子树是否为bst
		left = VerifySquenceOfBST(sequence[:i])
	}

	right := true
	if i < length-1 {
		// 判断右子树是否为bst
		right = VerifySquenceOfBST(sequence[i:length-1])
	}

	return left && right
}
