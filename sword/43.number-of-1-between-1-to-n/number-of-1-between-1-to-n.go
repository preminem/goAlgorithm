package problem43

/*
题：
输入一个整数n，求1~n这n个整数的十进制表示中1出现的次数。

分析：
记每一位的权值为base，位值为weight，该位之前的数是round，该位之后的数是former
则
若weight为0，则1出现次数为round*base
若weight为1，则1出现次数为round*base+former+1
若weight大于1，则1出现次数为round*base+base
*/

func numberOf1Between1AndN(n int) int {
	count := 0
	var round, former, weight int
	for base := 1; base <= n; base *= 10 {
		round = n / (base * 10)
		former = n % base
		weight = (n / base) % 10
		if 0 == weight {
			count += base * round
		} else if 1 == weight {
			count += base*round + former + 1
		} else {
			count += (round + 1) * base
		}
	}
	return count
}
