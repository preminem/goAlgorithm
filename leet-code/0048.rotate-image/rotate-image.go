package problem48

/*
题目
You are given an n x n 2D matrix representing an image.
Rotate the image by 90 degrees (clockwise).
Follow up: Could you do this in-place?

解题思路
对应的4个点，依次互换位置。
*/

func rotate(m [][]int) {
	n := len(m)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			temp := m[i][j]
			// 左边的行 等于 右边的列
			m[i][j] = m[n-j-1][i]
			m[n-j-1][i] = m[n-i-1][n-j-1]
			m[n-i-1][n-j-1] = m[j][n-i-1]
			m[j][n-i-1] = temp
		}
	}
}
