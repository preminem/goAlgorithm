package main

//container/list包实现了基本的双向链表功能，包括元素的插入、删除、移动功能
import (
	"container/list"
	"fmt"
	"strings"
)

type MyStack struct {
	List *list.List
}

type MyQueue struct {
	List *list.List
}

type BinaryTree struct {
	Value interface{}
	Left  *BinaryTree
	Right *BinaryTree
}

type Tree struct {
	Value    interface{}
	Children []*Tree
}

func (stack *MyStack) pop() interface{} {
	if elem := stack.List.Back(); elem != nil {
		stack.List.Remove(elem)
		return elem.Value
	}
	return nil
}

func (stack *MyStack) push(elem interface{}) {
	stack.List.PushBack(elem)
}

func (queue *MyQueue) pop() interface{} {
	if elem := queue.List.Front(); elem != nil {
		queue.List.Remove(elem)
		return elem.Value
	}
	return nil
}

func (queue *MyQueue) push(elem interface{}) {
	queue.List.PushBack(elem)
}

func preOrderRecur(node *BinaryTree) {
	if node == nil {
		return
	}

	fmt.Println(node.Value)
	preOrderRecur(node.Left)
	preOrderRecur(node.Right)
}

func inOrderRecu(node *BinaryTree) {
	if node == nil {
		return
	}

	inOrderRecu(node.Left)
	fmt.Println(node.Value)
	inOrderRecu(node.Right)
}

func posOrderRecu(node *BinaryTree) {
	if node == nil {
		return
	}

	posOrderRecu(node.Left)
	posOrderRecu(node.Right)
	fmt.Println(node.Value)
}

func preOrder(node *BinaryTree) {
	stack := MyStack{List: list.New()}
	stack.push(node)

	elem := stack.pop()
	for elem != nil {
		node, _ := elem.(*BinaryTree)

		fmt.Println(node.Value)
		if right := node.Right; right != nil {
			stack.push(right)
		}
		if left := node.Left; left != nil {
			stack.push(left)
		}

		elem = stack.pop()
	}
}

func inOrder(node *BinaryTree) {
	stack := MyStack{List: list.New()}
	current := node

	for stack.List.Len() > 0 || current != nil {
		if current != nil {
			stack.push(current)
			current = current.Left
		} else {
			current = stack.pop().(*BinaryTree)
			fmt.Println(current.Value)
			current = current.Right
		}
	}
}

func postOrder(node *BinaryTree) {
	stack1, stack2 := MyStack{List: list.New()}, MyStack{List: list.New()}
	stack1.push(node)

	for stack1.List.Len() > 0 {
		elem := stack1.pop().(*BinaryTree)
		stack2.push(elem)

		if elem.Left != nil {
			stack1.push(elem.Left)
		}

		if elem.Right != nil {
			stack1.push(elem.Right)
		}
	}

	for stack2.List.Len() > 0 {
		elem := stack2.pop().(*BinaryTree)
		fmt.Println(elem.Value)
	}
}

func levelOrder(node *BinaryTree) {
	var nlast *BinaryTree
	last := node
	level := 1
	queue := MyQueue{List: list.New()}

	fmt.Println(fmt.Sprintf("-----this is %d level-----", level))
	queue.push(node)
	for queue.List.Len() > 0 {
		node := queue.pop().(*BinaryTree)

		if node.Left != nil {
			queue.push(node.Left)
			nlast = node.Left
		}

		if node.Right != nil {
			queue.push(node.Right)
			nlast = node.Right
		}

		fmt.Println(node.Value)
		if last == node && (node.Left != nil || node.Right != nil) {
			last = nlast
			level++
			fmt.Println()
			fmt.Println(fmt.Sprintf("-----this is %d level-----", level))
		}
	}
}

func levelTreeOrder(node *Tree) {
	var nlast *Tree
	last := node
	queue := MyQueue{List: list.New()}
	queue.push(node)

	for queue.List.Len() > 0 {
		node := queue.pop().(*Tree)

		for _, elem := range node.Children {
			queue.push(elem)
			nlast = elem
		}

		fmt.Println(node.Value)

		if last == node {
			last = nlast
			fmt.Println()
		}

	}
}

func preOrderToStr(node *BinaryTree) (ret string) {
	if node == nil {
		return "#!"
	}

	ret += fmt.Sprintf("%d!", node.Value)
	ret += preOrderToStr(node.Left)
	ret += preOrderToStr(node.Right)
	return ret
}

func strToBinaryTree(arr []string, index *int) *BinaryTree {
	if *index >= len(arr) {
		return nil
	}

	if arr[*index] == "#" {
		*index++
		return nil
	}

	node := &BinaryTree{}
	node.Value = arr[*index]
	*index++

	node.Left = strToBinaryTree(arr, index)
	node.Right = strToBinaryTree(arr, index)
	return node
}
