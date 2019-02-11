package problem04

/*
题：
二维数组中的查找
在一个二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
请完成一个这样的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

分析： 从右上角开始，每次将搜索值与右上角的值比较，如果大于右上角的值，则直接去除1行，否则，则去掉1列。
如下图显示了查找13的轨迹。首先与右上角15比较，13<15，所以去掉最右1列，然后与11比较，这是13>11，去掉最上面1行…以此类推，最后找到13。
算法复杂度O(n)，最坏情况需要2n步，即从右上角开始查找，而要查找的目标值在左下角的时候。
*/

func find(matrix [][]int, rows, columns, number int) bool {
	if matrix != nil && rows > 0 && columns > 0 {
		row := 0
		column := columns - 1
		for row < rows && column >= 0 {
			//判断是否越界
			if matrix[row][column] == nil {
				return false
			} else if matrix[row][column] == number {
				return true
			} else if matrix[row][column] > number {
				column--
			} else {
				row++
			}
		}
	}
	return false
}
