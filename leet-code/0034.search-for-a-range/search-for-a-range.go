package problem0035

/*
Topic:
在排序数组中查找元素的第一个和最后一个位置
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
你的算法时间复杂度必须是 O(log n) 级别。
如果数组中不存在目标值，返回 [-1, -1]。

Example:
示例 1:
输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]

示例 2:
输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]
 */

func searchRange(nums []int, target int) []int {
	// 查看target是否存在与nums中
	index := search(nums, target)
	if index == -1 {
		return []int{-1, -1}
	}

	// 利用二分法，查找第一个target
	first := index
	for {
		f := search(nums[:first], target)
		if f == -1 {
			break
		}
		first = f
	}

	// 利用二分法，查找最后一个target
	last := index
	for {
		l := search(nums[last+1:], target)
		if l == -1 {
			break
		}
		// 注意这里与查找first的不同
		last += l + 1
	}

	return []int{first, last}
}

// 二分查找法
func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	var median int

	for low <= high {
		median = (low + high) / 2

		switch {
		case nums[median] < target:
			low = median + 1
		case nums[median] > target:
			high = median - 1
		default:
			return median
		}
	}
	return -1
}

