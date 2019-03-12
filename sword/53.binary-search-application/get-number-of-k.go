package problem53

/*
题：
统计一个数字在排序数组中出现的次数。

分析：
第一个k时，首先也是二分找到中间的k，然后和它上一个数字相比，如果还是k，则在前半段数组继续二分，直到找到中间的数为k且前一个不等于k。
用同样的方法，找最后一个k
时间复杂度为O(logn)
 */


func getNumberOfK(numbers []int, k int) int {
	if nil == numbers {
		return 0
	}
	first := getFirstK(numbers, k, 0, len(numbers)-1)
	last := getLastK(numbers, k, 0, len(numbers)-1)
	if first > -1 && last > -1 {
		return last - first + 1
	}
	return 0
}

func getFirstK(numbers []int, k, start, end int) int {
	if start > end {
		return -1
	}

	middleIndex := (start + end) / 2
	middleData := numbers[middleIndex]
	if middleData == k {
		if (0 < middleIndex && numbers[middleIndex-1] != k) || 0 == middleIndex {
			return middleIndex
		} else {
			end = middleIndex - 1
		}
	} else if middleData > k {
		end = middleIndex - 1
	} else {
		start = middleIndex + 1
	}
	return getFirstK(numbers, k, start, end)
}

func getLastK(numbers []int, k, start, end int) int {
	if start > end {
		return -1
	}

	middleIndex := (start + end) / 2
	middleData := numbers[middleIndex]
	if middleData == k {
		if (len(numbers)-1 > middleIndex && numbers[middleIndex+1] != k) || len(numbers)-1 == middleIndex {
			return middleIndex
		} else {
			start = middleIndex + 1
		}
	} else if middleData < k {
		start = middleIndex + 1
	} else {
		end = middleIndex - 1
	}
	return getLastK(numbers, k, start, end)
}