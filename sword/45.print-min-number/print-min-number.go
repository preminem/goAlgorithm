package problem45

/*
题：
输入一个正整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。例如，输入数组{3,32,321}，则打印出这
三个数字能排成的最小数字321323.

分析：
定义一种新的比较两个数字大小的规则
 */

import (
	"sort"
	"strconv"
)

type intSlice []int

func (p intSlice) Len() int {
	return len(p)
}

func (p intSlice) Less(i, j int) bool {
	return strconv.Itoa(p[i])+strconv.Itoa(p[j]) < strconv.Itoa(p[j])+strconv.Itoa(p[i])
}

func (p intSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func printMinNumber(numbers []int) string {
	if nil == numbers {
		return ""
	}

	sort.Sort(intSlice(numbers))
	var ans string
	for _, value := range numbers {
		ans += strconv.Itoa(value)
	}
	return ans
}
