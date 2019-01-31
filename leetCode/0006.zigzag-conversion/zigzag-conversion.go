package problem0006

import (
	"bytes"
)

/*
Topic:
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR" Write the code that will take a string and make this conversion given a number of rows:

func convert(text string, nRows int) string{...}
convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR".

Thought：
输入"ABCDEFGHIJKLMNOPQRSTUVWXYZ"和参数5后，得到答案"AGMSYBFHLNRTXZCEIKOQUWDJPV"， 按照题目的摆放方法，可得：

A   I   Q   Y
B  HJ  PR  XZ
C G K O S W
DF  LN  TV
E   M   U
可以看到，各行字符在原字符串中的索引号为

0行，0, 8, 16, 24
1行，1, 7,9, 15,17, 23,25
2行，2, 6, 10, 14, 18, 22
3行，3,5, 11,13, 19,21
4行，4, 12, 20

可以总结出以下规律
(1)第一行和最后一行下标间隔都是p=numRows×2-2 = 8 ;
(2)中间行的间隔是周期性的,第i行的间隔是: p–2*r,2*r,p–2*r,2*r,p–2*r,2*r,…

0行， 0×p，1×p，...
r行， r，1×p-r，1×p+r，2×p-r，2×p+r，...
最后一行， numRow-1, numRow-1+1×p，numRow-1+2×p，...
只需编程依次处理各行即可。

Complexity：
时间复杂度：O(n)。
空间复杂度O(n)。
*/

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	res := bytes.Buffer{}
	// p pace 步距
	p := numRows*2 - 2

	// 处理第一行
	for i := 0; i < len(s); i += p {
		res.WriteByte(s[i])
	}

	// 处理中间的行
	for r := 1; r <= numRows-2; r++ {
		// 添加r行的第一个字符
		res.WriteByte(s[r])

		for k := p; k-r < len(s); k += p {
			res.WriteByte(s[k-r])
			if k+r < len(s) {
				res.WriteByte(s[k+r])
			}
		}
	}

	// 处理最后一行
	for i := numRows - 1; i < len(s); i += p {
		res.WriteByte(s[i])
	}

	return res.String()
}
