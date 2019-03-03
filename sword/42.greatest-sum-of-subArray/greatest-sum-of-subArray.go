package problem42

/*
题：
输入一个正型数组，数组里有正数也有负数。数组中的一个或连续多个整数组成一个子数组。求所有子数组和的最大值。要求时间复杂度O(n).

分析：
当currentSum为负数时抛弃前面的累加，重新累加。当currentSum超过greatestSum时，更新greatestSum
*/

func findGreatestSumOfSubArray(data []int) (int, bool) {
	if nil == data || len(data) <= 0 {
		return 0, false
	}
	var currentSum, greatestSum int
	for i := 0; i < len(data); i++ {
		if currentSum <= 0 {
			currentSum = data[i]
		} else {
			currentSum += data[i]
		}
		if currentSum > greatestSum {
			greatestSum = currentSum
		}
	}
	return greatestSum, true
}
