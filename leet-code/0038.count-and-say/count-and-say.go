package problem0038

/*
Topic:
报数
报数序列是一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：
1.     1
2.     11
3.     21
4.     1211
5.     111221
1 被读作  "one 1"  ("一个一") , 即 11。
11 被读作 "two 1s" ("两个一"）, 即 21。
21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。

给定一个正整数 n（1 ≤ n ≤ 30），输出报数序列的第 n 项。

注意：整数顺序将表示为一个字符串。

Example:
示例 1:
输入: 1
输出: "1"

示例 2:
输入: 4
输出: "1211"

Thought:
利用byte记录上一次出现的次数和元素
*/

func countAndSay(n int) string {
	buf := []byte{'1'}

	for n > 1 {
		buf = say(buf)
		n--
	}

	return string(buf)
}

func say(buf []byte) []byte {
	// res 长度不会超过 buf 的两倍，所以，可以事先指定容量，加快append的速度
	res := make([]byte, 0, len(buf)*2)

	i, j := 0, 1
	for i < len(buf) {
		// 利用 j ，找到下一个不同的元素
		for j < len(buf) && buf[j] == buf[i] {
			j++
		}

		// res 中 res[i] 表示 res[i+1] 的个数，i 为0,2,4,6,...
		res = append(res, byte(j-i+'0'), buf[i])

		// 移动 i 到 j
		i = j
	}

	return res
}
