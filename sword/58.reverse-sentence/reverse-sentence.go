package problem58

/*
题目一：
输入一个英文句子，翻转句子中单词的顺序，但单词内字符的句子顺序不变。标点符号和普通字母一样处理
例如输入字符串“I am a student.”输出“student.  a am I”

分析：
1翻转句子中所有的字符——>.tueduts a ma I
2翻转每个单词中的字符顺序
*/

func reverseSentence(str string) string {
	if len(str) < 1 {
		return ""
	}

	copyCharArray := []rune(str)
	reverse(copyCharArray)

	var begin, end int
	for end <= len(copyCharArray) {
		if end == len(copyCharArray) {
			reverse(copyCharArray[begin:end])
			break
		}

		if copyCharArray[end] != ' ' {
			end++
		} else {
			reverse(copyCharArray[begin:end])
			end++
			begin = end
		}
	}
	return string(copyCharArray)
}

func reverse(str []rune) {
	copyStr := str

	begin, end := 0, len(copyStr)-1
	for begin < end {
		copyStr[begin], copyStr[end] = copyStr[end], copyStr[begin]
		begin++
		end--
	}
}

/*
题目二：
左旋转字符串
对于一个给定的字符序列S，请你把其循环左移K位后的序列输出。例如，字符序列S=”abcXYZdef”,要求输出循环左移3位后的结果，即“XYZdefabc”。

分析：
和上面的思路类似，我们可以将字符串分成两部分，先分别翻转两部分子字符串；
然后对翻转后的字符串整体再进行一次翻转。

*/
func leftRotateString(str string, n int) string {
	if len(str) < n {
		return ""
	}

	copyCharArray := []rune(str)
	reverse(copyCharArray[0:n])
	reverse(copyCharArray[n:])
	reverse(copyCharArray[:])

	return string(copyCharArray)
}
