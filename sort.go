package main

import (
	"fmt"
)

func main() {
	values := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	fmt.Println(values)
	// changeSort(values)
	// bubbleSort(values)
	// insertSort(values)
	// quickSort(values)
	// fmt.Println(values)
	// mergeSortEntrance(values)
	// fmt.Println(values)
	// shellSort(values)
	// fmt.Println(values)
	HeapSort(values)
	fmt.Println(values)
}

//选择排序，每次选择最小的，时间复杂度O(n^2)，稳定，空间复杂度O(1)
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

//冒泡排序，O(n^2)，稳定，空间复杂度O(1)
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

//插入排序，O(n^2)，稳定，空间复杂度O(1)
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

//快速排序On(nlog2n)，不稳定，空间复杂度O(log2n)~O(n)
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

//归并排序，时间复杂度O(nlog2n),稳定，空间复杂度O(n+log2n)≈O(n)
func mergeSortEntrance(values []int) {
	temp := make([]int, len(values))
	mergeSort(values, temp)
}

//递归的分解数列，再合并数列，就完成了归并排序
func mergeSort(values []int, temp []int) {
	if len(values) > 1 {
		middle := (len(values) - 1) / 2
		mergeSort(values[:middle+1], temp)
		mergeSort(values[middle+1:], temp)
		mergeArray(values, middle, temp)
	}
}

//将values[:middle+1]和values[middle+1:]两段有序数组借助temp辅助空间合并成完整有序数组，复杂度O(n)
func mergeArray(values []int, middle int, temp []int) {
	i, j, k := 0, middle+1, 0
	for i <= middle && j <= len(values)-1 {
		if values[i] <= values[j] {
			temp[k] = values[i]
			k++
			i++
		} else {
			temp[k] = values[j]
			k++
			j++
		}
	}
	for i <= middle {
		temp[k] = values[i]
		k++
		i++
	}
	for j <= len(values)-1 {
		temp[k] = values[j]
		k++
		j++
	}

	for ii := 0; ii < k; ii++ {
		values[ii] = temp[ii]
	}
}

//希尔排序，时间复杂度不好估计，不稳定，空间复杂度O(1)
func shellSort(values []int) {
	//初始化增量
	increment := len(values) + 1
	for increment > 1 {
		//每次减小增量
		increment = increment/3 + 1
		//对每个划分进行直接插入排序
		for i := increment; i < len(values); i++ {
			//一次插入
			if values[i] < values[i-increment] {
				temp := values[i]
				var j int
				for j = i - increment; j >= increment && values[j] > temp; j -= increment {
					values[j+increment] = values[j]
				}
				values[j+increment] = temp
			}
		}
		//直到增量为1
		if increment == 1 {
			break
		}
	}
}

//堆排序，时间复杂度O(nlog2n),不稳定，空间复杂度O(1)
func HeapSort(values []int) {
	//将无序序列构建成一个堆，升序选择大顶堆，降序选择小顶堆
	buildHeap(values)
	for i := len(values); i > 1; i-- {
		//将堆顶元素与末尾元素交换，最大元素沉到数组末端
		values[0], values[i-1] = values[i-1], values[0]
		//重新调整结构，使其满足堆定义，然后继续交换堆顶元素与当前末尾元素，反复执行调整+交换步骤，直到整个序列有序。
		adjustHeap(values[:i-1], 0)
	}
}

func buildHeap(values []int) {
	for i := len(values); i >= 0; i-- {
		adjustHeap(values, i)
	}
}

//调整pos位置的结点,pos与child的交换可能导致混乱，因此设置child为新的pos继续调整
func adjustHeap(values []int, pos int) {
	node := pos
	length := len(values)
	for node < length {
		var child int = 0
		if 2*node+2 < length {
			if values[2*node+1] > values[2*node+2] {
				child = 2*node + 1
			} else {
				child = 2*node + 2
			}
		} else if 2*node+1 < length {
			child = 2*node + 1
		}
		if child > 0 && values[child] > values[node] {
			values[node], values[child] = values[child], values[node]
			node = child
		} else {
			break
		}
	}
}
