package problem53

/*
题：
数组中数值和下表相等的元素
假设单调递增的数组里每个元素都是整数且都是唯一的。请实现函数，找出数组中任意一个数值等于其下标的元素。如
{-3,-1,1,3,5}，数字3和下标相等

分析：
若Index = value，则找到
若index > value，在右半段
若index < value，在左半段
 */

func getNumberSameAsIndex(numbers []int) int {
	if nil == numbers {
		return -1
	}

	left, right := 0, len(numbers)-1
	for left <= right {
		middle := (left + right) / 2
		if numbers[middle] == middle {
			return middle
		} else if numbers[middle] > middle {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return -1
}