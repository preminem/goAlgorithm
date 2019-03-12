package problem53

/*
题：
0~n-1中缺失的数字
一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在0~n-1之内。在范围0~n-1内的n个数字有且只有一个不在数组中，找出该数字

分析：
遇到有序数组中的查找，往往要想到二分，因为能把复杂度降到O(logn)
data[index_mid] = data_mid
1、若 index_mid = data_mid，则要找的数字在右半段
2、(a)若index_mid != data_mid且data[index_mid - 1] = index_mid - 1，也就是当前索引与值不等，但上一个索引与数值相等，则index_mid就是我们要找的数字
　 (b)若index_mid != data_mid且data[index_mid - 1] != index_mid - 1，那么我们要找的数字在左半段
 */

func getMissingNumber(numbers []int) int {
	if nil == numbers || len(numbers) <= 0 {
		return -1
	}
	left, right := 0, len(numbers)-1
	for left <= right {
		middle := (left + right) >> 1
		if numbers[middle] != middle {
			if 0 == middle || numbers[middle-1] == middle-1 {
				return middle
			}
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	if left == len(numbers) {
		return len(numbers)
	}
	return -1
}