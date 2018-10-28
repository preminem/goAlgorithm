package main

import (
	"fmt"
)

const MAX = 10

func main() {
	values := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	fmt.Println(values)
	// changeSort(values)
	// bubbleSort(values)
	// insertSort(values)
	quickSort(values)
	fmt.Println(values)
}

//选择排序，每次选择最小的，O(n^2)
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

//冒泡排序，O(n^2)
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

//插入排序，O(n^2)
//i代表待插入的数字，j代表排好序并且待移动的数字
func insertSort(values []int) {
	for i := 1; i < len(values); i++ {
		if values[i] < values[i-1] {
			temp := values[i]
			var j int
			for j = i - 1; j >= 0 && values[j] > temp; j-- {
				values[j+1] = values[j]
			}
			values[j+1] = temp
		}
	}
	fmt.Println(values)
}

//快速排序On(nlogn)，不稳定
//取左边第一个数做为基数，j先从后往前扫描，i后从前往后扫描，那么相遇的数一定比基数小
func quickSort(values []int) {
	left, right := 0, len(values)-1
	if left < right {
		i, j := left, right
		for {
			for i < j && values[j] >= values[left] {
				j--
			}
			for i < j && values[i] <= values[left] {
				i++
			}
			if i >= j {
				break
			}
			values[i], values[j] = values[j], values[i]
		}
		values[i], values[left] = values[left], values[i]
		quickSort(values[:i])
		quickSort(values[i+1:])
	}
}
