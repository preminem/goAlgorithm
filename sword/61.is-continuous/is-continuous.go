package problem61

/*
题目：扑克牌中的顺子
从扑克牌中随机抽取5张牌，判断是不是一个顺子，即这5张牌是不是连续的。2~10为数字本身，A为1，J为11，Q为12，K为13，而大、小王可以看成任意数字。

分析：
分步实现
通过对问题抽象，我们我们可以看出，题目就是一个判断数组连续，并且其中将大小王作为0，可以用0补充不连续位置的题目；
首先把数组排序，虽然对于常规的排序，一般需要O(nlogn)时间复杂度，但是在这种小问题中，我们可以直接用内置的算法模块，小数据的时间复杂度在这里不重要；
统计数组中0的个数；这里可以设置一个变量即为king的个数，如果对于其他多王的情况，可以通过修改实现扩展；
统计排序之后的数组中相邻数字之间的空缺总数；
对0的个数和空缺数进行比较，如果足够填补，那么即连续；否则不连续；
同时需要注意：如果数组中的非0数字重复出现，则该数组不是连续的
*/

import "sort"

func isContinuous(numbers []int) bool {
	if nil == numbers || len(numbers) < 5 {
		return false
	}

	sort.Ints(numbers)
	numberOfZero, numberOfGap := 0, 0
	for _, value := range numbers {
		if value == 0 {
			numberOfZero++
		}
	}

	small, big := 0, 1
	for big < len(numbers) {
		if numbers[small] == numbers[big] {
			return false
		}
		numberOfGap += numbers[big] - numbers[small] - 1
		small = big
		big++
	}
	if numberOfGap > numberOfZero {
		return false
	}
	return true
}
