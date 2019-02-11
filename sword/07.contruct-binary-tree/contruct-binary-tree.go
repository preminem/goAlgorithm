package problem07

/*
题：
重建二叉树
输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。 假设输入的前序遍历和中序遍历都不含重复数字。
例如，输出前序遍历序列{1, 2, 4, 7, 3, 5, 6, 8} 和中序遍历序列{4, 7, 2, 1, 5, 3, 8, 6}

分析: 前序遍历输出: 1, 2, 4, 7, 3, 5, 6, 8 后续遍历输出: 4, 7, 2, 1, 5, 3, 8, 6
前序遍历第一个值即为根节点。 中序遍历中找到根节点位置，根据中序遍历特点，1左边的在根节点的左子树，右边的为根节点的右子树
同样先序遍历中，根节点后面的三个数字就是3个左子树节点的值，再后面的所有数字都是右子树节点的值。
分析到这里，我们即可想到可以使用递归完成重建二叉树。分别对左子树右子树进行递归操作
*/
import (
    "fmt"
	"github.com/preminem/goAlgorithm/kit"
)

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }
type TreeNode = kit.TreeNode

func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}

	msg := fmt.Sprintf("%d 不存在于 %v 中", val, nums)
	panic(msg)
}

// PreIn2Tree 把 preorder 和 inorder 切片转换成 二叉树
func PreIn2Tree(pre, in []int) *TreeNode {
	if len(pre) != len(in) {
		panic("preIn2Tree 中两个切片的长度不相等")
	}

	if len(in) == 0 {
		return nil
	}

	res := &TreeNode{
		Val: pre[0],
	}

	if len(in) == 1 {
		return res
	}

	idx := indexOf(res.Val, in)

	res.Left = PreIn2Tree(pre[1:idx+1], in[:idx])
	res.Right = PreIn2Tree(pre[idx+1:], in[idx+1:])

	return res
}
