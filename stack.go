package goAlgorithm

import (
	"errors"
	"reflect"
)

// 栈定义
type Stack struct {
	values    []interface{}
	valueType reflect.Type
}

// 构造栈
func NewStack(valueType reflect.Type) *Stack {
	return &Stack{values: make([]interface{}, 0), valueType: valueType}
}

// 判断值是否符合栈类型
func (stack *Stack) isAcceptableValue(value interface{}) bool {
	if value == nil || reflect.TypeOf(value) != stack.valueType {
		return false
	}
	return true
}

// 入栈
func (stack *Stack) Push(v interface{}) bool {
	if !stack.isAcceptableValue(v) {
		return false
	}
	stack.values = append(stack.values, v)
	return true
}

// 出栈
func (stack *Stack) Pop() (interface{}, error) {
	if stack == nil || len(stack.values) == 0 {
		return nil, errors.New("stack empty")
	}
	v := stack.values[len(stack.values)-1]
	stack.values = stack.values[:len(stack.values)-1]
	return v, nil
}

// 获取栈顶元素
func (stack *Stack) Top() (interface{}, error) {
	if stack == nil || len(stack.values) == 0 {
		return nil, errors.New("stack empty")
	}
	return stack.values[len(stack.values)-1], nil
}

// 获取栈内元素个数
func (stack *Stack) Len() int {
	return len(stack.values)
}

// 判断栈是否为空
func (stack *Stack) Empty() bool {
	if stack == nil || len(stack.values) == 0 {
		return true
	}
	return false
}

// 获取栈内元素类型
func (stack *Stack) ValueType() reflect.Type {
	return stack.valueType
}
