package problem03

/*
题：
不修改数组找出重复的数字。
在一个长度为n+1的数组里的所有数字都在1~n的范围内。数组中某些数字是重复的，但是不知道有几个数字重复了，也不知道每个数字重复了几次。
请找出数组中任意一个重复的数字, 但是不能修改数组。
例如，如果输入长度为8的数组{2，3，5，4，3，2，6，7}，那么对应的输出是重复的数字2或者3。

方法一：创建一个长度为n+1辅助数组，如果原数组中被复制的是m，则把它复制到辅助数组中下表为m的位置。这样很容易发现哪个数字是重复的.
空间复杂度O(n)。
方法二：假设没有重复数字，那么在1~n的范围内只有n个数字，超过则有重复。我们把1~n的数字从中间数字m分为两部分，前面一半为1~m后一半为m+1~n。
如果1~m的数字超过了m，那么这一半的数字中一定包含了重复数字，否则重复数字在另一半中。继续把包含重复数字的区间一分为二，直到找到重复数字。
时间复杂度O(nlogn)。这种算法也不能找出所有重复数字。
*/

func countRange(numbers []int, start, end int) int {
	count := 0
	for _, num := range numbers {
		if num >= start && num <= end {
			count++
		}
	}
	return count
}

func findDuplicationNumberInArrayNotEdit(numbers []int) int {
	length := len(numbers)
	if length <= 0 {
		return -1
	}

	for i := 0; i < length; i++ {
		if numbers[i] < 0 || numbers[i] > length-1 {
			return -1
		}
	}

	start := 1
	end := length - 1
	for start <= end {
		middle := ((end - start) >> 1) + start
		count := countRange(numbers, start, middle)
		if end == start {
			if count > 1 {
				return start
			} else {
				break
			}
		}

		if count > (middle - start + 1) {
			end = middle
		} else {
			start = middle + 1
		}
	}
	return -1
}
