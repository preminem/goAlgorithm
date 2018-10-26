package main

import (
	"fmt"
)

func main() {
	values := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	fmt.Println(values)
	changeSort(values)
	bubbleSort(values)
	insertSort(values)
}

//选择排序，每次选择最小的，On(n^2)
//i代表要确定数字的位置，j来寻找最小的数字和i交换
func changeSort(values []int) {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
	fmt.Println(values)
}

//冒泡排序，On(n^2)
//i代表冒泡的轮数和冒出来泡的个数，j代表从第0位开始冒泡交换
func bubbleSort(values []int) {
	for i := 0; i < len(values)-1; i++ {
		for j := 0; j < len(values)-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
			}
		}
	}
	fmt.Println(values)
}

func insertSort(values []int) {
	for i := 1; i < len(values); i++ {
		for j := 0; j < i; j++ {
			if values[j] < values[i] {

			} else {
				continue
			}
		}
	}
	fmt.Println(values)
}
