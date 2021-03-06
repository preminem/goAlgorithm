package search

//顺序查找O(n)
func SequenceSearch(values []int, value int) int {
	for i := 0; i < len(values); i++ {
		if values[i] == value {
			return i
		}
	}
	return -1
}

//二分查找O(log2n),适用于有序非动态序列
func BinarySearch(values []int, value int) int {
	low, high := 0, len(values)-1
	for low <= high {
		mid := low + (high-low)/2
		if values[mid] == value {
			return mid
		} else if values[mid] > value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

//差值查找，O(log2(log2n))，基于二分查找算法，将查找点的选择改进为自适应选择，提高查找效率
func InsertSearch(values []int, value int) int {
	low, high := 0, len(values)-1
	if high == 0 && values[0] == value {
		return 0
	}
	//low不能等于high
	for low < high {
		//根据 (查找值 - 起始值) / (末位值 - 起始值) 的比例来决定中间值的下标
		mid := low + (high-low)*(value-values[low])/(values[high]-values[low])
		if values[mid] == value {
			return mid
		} else if values[mid] > value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
