package problem63

/*
题目：股票的最大利润（一次卖出）
假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可获得的最大利润是多少？

思路标签：
记录当前最小值和最大差值

解答：
最大利润无外乎就是计算后面的数字减去前面的数字得到的一个最大的差值；
求总体的最大差值，需要的数据：当前的【最小值】，当前的【最大差值】；遍历求解即可
*/

func maxDiff(numbers []int) int {
	if nil == numbers || len(numbers) < 2 {
		return -1
	}

	min := numbers[0]
	max := numbers[1] - min
	for i := 2; i < len(numbers); i++ {
		if numbers[i-1] < min {
			min = numbers[i-1]
		}

		currentDiff := numbers[i] - min
		if currentDiff > max {
			max = currentDiff
		}
	}
	return max
}
