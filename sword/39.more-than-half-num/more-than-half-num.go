package problem39

/*
题：
数组中出现次数超过一半的数字
数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
例如，输入一个长度为9的数组{1,2,3,2,2,2,5,4,2}。由于数字2在数组中出现了5次，超过数组长度的一半，因此输出2.

分析：
方法一：先排序，再统计。复杂度O(nlogn)
方法二：排序之后位于数组中间的数字一定就是那个出现次数超过一半的数字。（统计学中的中位数）
受快速排序启发，每过一轮判断选中的数字>n/2还是<n/2。复杂度O(n)，会改变数组顺序。
方法三：根据数组特点找出时间复杂度为O(n)的解法。
 */

 //因为要找的数字比其他数字出现的次数总和还要多，那么要找的数字肯定是最后一次把次数设为1时对应的数字
func moreThanHalfNum(numbers []int) int {
	if nil == numbers || 0 >= len(numbers) {
		return 0
	}
	ans := numbers[0]
	times := 1
	for i := 1; i < len(numbers); i++ {
		if 0 == times {
			ans = numbers[i]
			times = 1
		} else if ans == numbers[i] {
			times++
		} else {
			times--
		}
	}
	return ans
}
