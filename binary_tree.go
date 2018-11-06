package goAlgorithm

import (
	"fmt"
	"reflect"
)

// 二叉树定义
type BinaryTree struct {
	Data   interface{}
	Lchild *BinaryTree
	Rchild *BinaryTree
}

// 构造方法
func NewBinaryTree(data interface{}) *BinaryTree {
	return &BinaryTree{Data: data}
}

//先序遍历，递归
func (bt *BinaryTree) PreTraverse(t *TreeNode) {
	if t != nil {
		fmt.Printf("%d/", t.Value)
		PreTraverse(t.Left)
		PreTraverse(t.Right)
	}
}

// 先序遍历，非递归
func (bt *BinaryTree) PreStackTraverse() []interface{} {
	t := bt
	stack := NewStack(reflect.TypeOf(bt))
	res := make([]interface{}, 0)
	for t != nil || !stack.Empty() {
		for t != nil {
			res = append(res, t.Data)
			stack.Push(t)
			t = t.Lchild
		}
		if !stack.Empty() {
			v, _ := stack.Pop()
			t = v.(*BinaryTree)
			t = t.Rchild
		}
	}
	return res
}

//中序遍历，递归
func (bt *BinaryTree) MidTraverse(t *TreeNode) {
	if t != nil {
		MidTraverse(t.Left)
		fmt.Printf("%d/", t.Value)
		MidTraverse(t.Right)
	}
}

// 中序遍历
func (bt *BinaryTree) MidStackTraverse() []interface{} {
	t := bt
	stack := NewStack(reflect.TypeOf(bt))
	res := make([]interface{}, 0)
	for t != nil || !stack.Empty() {
		for t != nil {
			stack.Push(t)
			t = t.Lchild
		}
		if !stack.Empty() {
			v, _ := stack.Pop()
			t = v.(*BinaryTree)
			res = append(res, t.Data)
			t = t.Rchild
		}
	}
	return res
}

//后序遍历，递归
func (bt *BinaryTree) PostTraverse(t *TreeNode) {
	if t != nil {
		PostTraverse(t.Left)
		PostTraverse(t.Right)
		fmt.Printf("%d/", t.Value)
	}
}

// 后续遍历，非递归
func (bt *BinaryTree) PostStackTraverse() []interface{} {
	t := bt
	stack := NewStack(reflect.TypeOf(bt))
	s := NewStack(reflect.TypeOf(true))
	res := make([]interface{}, 0)
	for t != nil || !stack.Empty() {
		for t != nil {
			stack.Push(t)
			s.Push(false)
			t = t.Lchild
		}
		for flag, _ := s.Top(); !stack.Empty() && flag.(bool); {
			s.Pop()
			v, _ := stack.Pop()
			res = append(res, v.(*BinaryTree).Data)
			flag, _ = s.Top()
		}
		if !stack.Empty() {
			s.Pop()
			s.Push(true)
			v, _ := stack.Top()
			t = v.(*BinaryTree)
			t = t.Rchild
		}
	}
	return res
}
