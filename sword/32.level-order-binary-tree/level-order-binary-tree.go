package problem32

/*
题一：
从上倒下打印二叉树
从上往下打印出二叉树的每个结点，同一层的结点按照从左到右的顺序打印。
 */
import (
	"github.com/preminem/goAlgorithm/data-structure/binary-tree"
	"container/list"
	"fmt"
)

type BinaryTree = binary_tree.BinaryTree
type MyQueue = binary_tree.MyQueue

func levelOrder(node *BinaryTree) {
	queue := MyQueue{List: list.New()}

	queue.push(node)
	for queue.List.Len() > 0 {
		node := queue.pop().(*BinaryTree)

		if node.Left != nil {
			queue.push(node.Left)
		}

		if node.Right != nil {
			queue.push(node.Right)
		}

		fmt.Println(node.Value)
	}
}

/*
题二：
分行从上到下打印二叉树
从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印一行。

分析：
与上一道层次遍历类似，但是另外需要两个变量：
1.toBePrint：当前层中还没打印的节点数
2.nextLevel：下一层节点的数目
 */
func printLevelOrderByRow(node *BinaryTree) {
	queue := MyQueue{List: list.New()}
	queue.push(node)

	nextLevel := 0
	toBePrint := 1

	for queue.List.Len() > 0 {
		currentNode := queue.pop().(*BinaryTree)
		fmt.Printf("%d  ", currentNode.Value)
		toBePrint--
		if currentNode.Left != nil {
			queue.push(currentNode.Left)
			nextLevel++
		}
		if currentNode.Right != nil {
			queue.push(currentNode.Right)
			nextLevel++
		}

		if toBePrint == 0 {
			fmt.Printf("\n")
			toBePrint = nextLevel
			nextLevel = 0
		}
	}
}

/*
题三：
之字形上到下打印二叉树
实现一个函数按照之字形顺序打印二叉树，即第一层从左到右的顺序打印，第二层从右到左的顺序打印，第三层再按照从左到右的顺序打印。

分析：
可以使用栈的特性来实现这个功能。
定义两个栈，s1,s2。分别用来存当前层和下一层
 */

type MyStack = binary_tree.MyStack

func PrintLevelOrderZigzag(node  *BinaryTree) {
	stacks := [2]*MyStack{
		&MyStack{List: list.New()},
		&MyStack{List: list.New()},
	}

	current := 0
	next := 1

	stacks[current].push(node)

	for stacks[0].List.Len() > 0 || stacks[1].List.Len() > 0 {
		currentNode := stacks[current].pop().(*BinaryTree)
		fmt.Printf("%d  ", currentNode.Value)
		//如果打印奇数层，则先保存左子节点再保存右子结点
		if current == 0 {
			if currentNode.Left != nil {
				stacks[next].push(currentNode.Left)
			}

			if currentNode.Right != nil {
				stacks[next].push(currentNode.Right)
			}
		////如果打印偶数层，则先保存右子节点再保存左子结点
		} else {
			if currentNode.Right != nil {
				stacks[next].push(currentNode.Right)
			}

			if currentNode.Left != nil {
				stacks[next].push(currentNode.Left)
			}

		}

		if stacks[current].List.Len() == 0 {
			fmt.Printf("\n")
			current = 1 - current
			next = 1 - next
		}
	}
}