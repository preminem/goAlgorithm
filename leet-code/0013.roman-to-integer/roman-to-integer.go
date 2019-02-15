package problem0013

/*
Topic:
Given a roman numeral, convert it to an integer.
Input is guaranteed to be within the range from 1 to 3999.

Thought:
此题，最关键的信息是
右加左减，左减数字必须为一位，比如8写成VIII，而非IIX。
解题思路
1.从右往左处理字符串。
2.当前字符代表的数字，小于右边字符的时候，总体减去当前字符代表的数字。
3.否则，总体加上当前字符代表的数字。
*/
func romanToInt(s string) int {
	res := 0
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	last := 0
	for i := len(s) - 1; i >= 0; i-- {
		temp := m[s[i]]

		sign := 1
		if temp < last {
			//小数在大数的左边，要减去小数
			sign = -1
		}

		res += sign * temp

		last = temp
	}

	return res
}