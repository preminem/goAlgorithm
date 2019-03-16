package problem62

/*
题目：
0、1…n-1这n个数字排成一个圆圈，从数字0开始，每次从这个圆圈里删除第m个数字。求出这个圆圈里剩下的最后一个数字。

例子：
如0、1、2、3、4这5个数字组成一个圆圈，从数字0开始每次删除第3个数字，则删除的前4个数字依次是2、0、4、1，因此最后剩下的数字是3.

分析：
f(n,m)表示每次在n个数字0,1,...,n-1中删除第m个数字最后剩下的数字。
f(n,m)= 0                 n=1
        [f(n-1,m)+m]%n    n>1
*/
func lastRemaining(n, m int) int {
	if n < 1 || m < 1 {
		return -1
	}

	last := 0
	for i := 2; i <= n; i++ {
		last = (last + m) % i
	}
	return last
}
