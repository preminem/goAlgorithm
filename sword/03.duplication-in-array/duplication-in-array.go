package problem03

/*
题：
找出数组中重复的数字
在一个长度为n的数组里的所有数字都在0~n-1的范围内。数组中某些数字是重复的，但是不知道有几个数字重复了，也不知道每个数字重复了几次。
请找出数组中任意一个重复的数字。例如，如果输入长度为7的数组{2，3，1，0，2，5，3}，那么对应的输出是重复的数字2或者3。

方法一：先排序，从排序的数组中找出重复的数字是一件很容易的事。时间复杂度O(nlogn)
方法二：申请一个哈希表。遍历数组时将哈希表没有的数字加入哈希表。时间复杂度O(n)，空间复杂度O(n)
方法三：边排序边搜索比较。从头到尾一次扫描这个数组中的每个数字。当扫描到下表为i的数字时，首先比较这个数字(用m表示)是不是等于i。
如果是，则接着扫描下一个数字；如果不是，则再拿它和第m个数字进行比较。如果它和第m个数字相等，就找到了一个重复的数字；
如果它和第m个数字不相等，就把第i个数字和第m个数字交换，把m放到属于他的位置。重复这个过程，直到发现一个重复的数字。
*/
import (
	"fmt"
)

func findDuplicationNumberInArray(numbers []int) int {
	length := len(numbers)
	if length <= 0 {
		return -1
	}

	for i := 0; i < length; i++ {
		if numbers[i] < 0 || numbers[i] > length-1 {
			return -1
		}
	}

	for i := 0; i < length; i++ {
		for i != numbers[i] {
			if numbers[i] == numbers[numbers[i]] {
				return numbers[i]
			}
			numbers[i], numbers[numbers[i]] = numbers[numbers[i]], numbers[i]
		}
	}

	return -1
}
