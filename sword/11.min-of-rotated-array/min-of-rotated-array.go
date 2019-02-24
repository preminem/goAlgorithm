package problem11

/*
题：
旋转数组的最小数字
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。
例如，数组{3, 4, 5, 1, 2}为{1, 2, 3, 4, 5}的一个旋转，该数组最小值为1

分析：
方法一：从头到尾遍历一边数组即可，时间复杂度O(n)
方法二: 旋转之后的数组分为两个子数组，而且前面子数组的元素都大于后面子数组的元素，最小的元素刚好是两个子数组的分界线，可以使用二分查找。
我们使用两个指针分别指向数组的第一个和最后一个元素，接着我们找到数组中间的元素，如果它处于前面的数组，那么它应该大于或者等于第一个指针指向的元素。
此时，数组中最小的元素应该位于该中间元素的后面。把第一个指针指向中间元素。
反之，它要是小于等于第二个指针所指元素，那么最小元素在中间元素的前面，把第二个指针移动到中间元素
按照上面思路，第一个指针总是指向前面递增数组的元素，第二个指针总是指向第二个递增数组的元素。
最终，第一个指针指向第一个数组的最后一个元素，第二个指针指向第二个数组的第一个元素。
也就是它们会指向两个相邻的元素，第二个指针指向的是数组中最小元素，这也是循环结束的条件。

在考虑一种特殊的情况: 原数组{0, 1, 1, 1, 1} 旋转后{1, 1, 1, 0, 1}。
此时两个指针指向的元素与中间数都等于1，此时无法判断最小元素所在哪个子数组，这种情况只能顺序查找。
*/

// 如果数组中的数字没有重复的话
//func indexOfMin(nums []int) int {
//	size := len(nums)
//	left, right := 0, size-1
//	for left < right {
//		mid := (left + right) / 2
//		if nums[right] < nums[mid] {
//			left = mid + 1
//		} else {
//			right = mid
//		}
//	}
//	return left
//}

func Min(numbers []int) int {
	length := len(numbers)
	index1, index2, indexMid := 0, length-1, 0
	for numbers[index1] >= numbers[index2] {
		if index2-index1 == 1 {
			indexMid = index2
			break
		}

		indexMid = (index1 + index2) / 2
		if numbers[index1] == numbers[index2] && numbers[indexMid] == numbers[index1] {
			return MinInOrder(numbers, index1, index2)
		}

		if numbers[indexMid] >= numbers[index1] {
			// 中间数大于第一个指针，说明在后面的区间
			index1 = indexMid
		} else if numbers[indexMid] <= numbers[index2] {
			index2 = indexMid
		}
	}
	return numbers[indexMid]
}

// 顺序遍历区间，返回最小值
func MinInOrder(numbers []int, index1, index2 int) int {
	result := numbers[index1]
	for i := index1 + 1; i < index2; i++ {
		if result > numbers[i] {
			result = numbers[i]
		}
	}
	return result
}
