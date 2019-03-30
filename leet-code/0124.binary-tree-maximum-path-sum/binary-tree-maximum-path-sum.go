package problem0124

/*
题：
二叉树中的最大路径和
给定一个非空二叉树，返回其最大路径和。
本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
*/
import (
	"github.com/preminem/goAlgorithm/kit"
)

type TreeNode = kit.TreeNode

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxSum := root.Val

	var dfs func(*TreeNode) int
	// 返回，以 root 为起点，所有可能路径的 sum 值中的最大值。
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := max(0, dfs(root.Left))
		right := max(0, dfs(root.Right))
		sum := left + root.Val + right
		if maxSum < sum {
			maxSum = sum
		}

		return max(left, right) + root.Val
	}

	dfs(root)

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
