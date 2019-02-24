package problem14

import "math"

/*
题：
剪绳子（动态规划与贪婪算法）
给你根长度为n的绳子,请把绳子剪成m段(m,n都是整数并且m,n>1)每段绳子长度记为k[0],k[1]...k[m],请问k[0]k[1]...*k[m]的最大乘积是多少？
例如,当n=8时,剪成2,3,3的三段,得到最大乘积18

动态规划: 如果是求最优解,通常是最大值最小值,而且该问题能够分解成若干个子问题,并且子问题间还有重叠的更小的子问题.就可以考虑动态规划来解决.
如本例,假设乘积最大值为f(n), 假设我们第一刀剪在长度为i(0<i<n)的位置, 于是把绳子剪成了i和n-i两段, 我们要得到问题的最优解f(n),
那么要用最优的方法把i和n-i的两段分别剪成若干段,使它们各自剪出的绳子乘积最大.也就是说整体问题的最优解是依赖与各个子问题的最优解.

复杂度：
动态规划：时间复杂度O(n^2) 空间复杂度O(n)
贪婪算法：时间复杂度O(1) 空间复杂度O(1)
*/

// 动态规划, 把切成m的任务分成两段.
func MaxProductAfterCyttingSolution1(length int) int {
	if length < 2 {
		return 0
	}
	if length == 2 {
		return 1
	}
	if length == 3 {
		return 2
	}

	//最优解存储在products里
	products := make([]int, length+1)
	products[0] = 0
	products[1] = 1
	products[2] = 2
	products[3] = 3

	max := 0

	// 算出f(i)的最优解
	for i := 4; i <= length; i++ {
		max = 0
		//在求f(i)之前每一个f(j)都被求出来了，为了求f(i)我们需要求出所有可能的f(j)*f(i-j)
		for j := 1; j <= i/2; j++ {
			product := products[j] * products[i-j]
			if max < product {
				max = product
			}

			products[i] = max
		}
	}

	max = products[length]
	return max
}

// 贪婪算法,当n>=5时，尽可能多地剪长度为3的绳子。当剩下的绳子长度为4时，把绳子剪成2段长度为2的绳子
func MaxProductAfterCyttingSolution2(length int) int {
	if length < 2 {
		return 0
	}

	if length == 2 {
		return 1
	}

	if length == 3 {
		return 2
	}

	// 尽可能减去长度为3的绳子段
	timesOf3 := length / 3

	// 当绳子最后剩下长度为4的时候,不能在剪去长度为3的绳子段
	// 此时更好的方法是把绳子剪成长度为2的两段
	if length-timesOf3*3 == 1 {
		timesOf3 -= 1
	}

	timesOf2 := (length - timesOf3*3) / 2

	return int(math.Pow(3, float64(timesOf3)) * math.Pow(2, float64(timesOf2)))
}
