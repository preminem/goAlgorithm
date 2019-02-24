package problem10

/*
题：
求斐波那契数列第n项

分析：
递归不是最佳解法，效率及其低下。
最佳解决办法很简单，用f(0)f(1)算出f(2)，在f(1)f(2)算出f(3)..依次类推算出f(n)

注：青蛙跳台阶问题也属于斐波那契数列的应用。
*/

func fib(num int) int {
	if num == 0 {
		return 0
	} else if num == 1 {
		return 1
	}

	fibNMinusOne := 1
	fibNMinusTwo := 0
	fibN := 0
	for i := 2; i <= num; i++ {
		fibN = fibNMinusOne + fibNMinusTwo
		fibNMinusTwo = fibNMinusOne
		fibNMinusOne = fibN
	}
	return fibN
}
