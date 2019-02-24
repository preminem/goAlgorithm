package problem15

/*
题：
请实现一个函数,输入一个整数,输出该数二进制表示中的1的个数. 例如,把9表示成二进制是1001,有2位是1.因此,如果输入9,则该函数输出2.
*/

//先把输入的整数num与1做&运算,判断num的最低位是否为1.接着把1左移一位,在和num做&预算,判断次低位是否为一.循环的次数为整数二进制的位数.
func NumberOf1Solution1(num int) int {
	count := 0
	flag := 1
	for flag > 0 {
		if num&flag != 0 {
			count++
		}
		flag = flag << 1
	}
	return count
}

//把一个整数减去1，再和原整数做与运算，会把该整数最右边的1变成0. 那么一个整数的二进制表示中有多少个1，就可以进行多少次这样的运算。
func NumberOf1Solution2(num int) int {
	count := 0

	for num != 0 {
		count++
		num = (num - 1) & num
	}
	return count
}
