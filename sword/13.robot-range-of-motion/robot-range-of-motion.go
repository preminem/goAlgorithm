package problem13

/*
题：
机器人的运动范围
地上有一个m行n列的方格。一个机器人从坐标（0，0）的格子开始移动，每次可以向左右上下移动一格，但不能进入行坐标和列坐标的数位之和大于k的格子。
例如，当k为18时，机器人能够进入方格(35,37)因为3+5+3+7=18但是不能进入(35,38)因为3+5+3+8=19。请问机器人能够到达多少个格子?

思路：
和上一题矩阵中的路径思路一样，起始位置固定为(0,0)，如果能进入(i,j)的格子，则再判断是否能进入(i, j-1) (i-1, j) (i+1, j), (i, j+1)的格子
*/

func movingCount(k, rows, cols int) int {
	startRow := 0
	startCol := 0
	visited := make([]bool, rows*cols) // 标记是否访问过

	// 回溯法求解
	count := movingCountCore(k, rows, cols, startRow, startCol, visited)

	return count
}

func movingCountCore(k, rows, cols, row, col int, visited []bool) int {
	count := 0
	if check(k, rows, cols, row, col, visited) {
		visited[row*cols+col] = true

		// 进入下一歌
		count = 1 + movingCountCore(k, rows, cols, row-1, col, visited) + movingCountCore(k, rows, cols, row, col-1, visited) + movingCountCore(k, rows, cols, row+1, col, visited) + movingCountCore(k, rows, cols, row, col+1, visited)
	}

	return count
}

// 判断是否能进入下一格
func check(k, rows, cols, row, col int, visited []bool) bool {
	if row >= 0 && row < rows && col >= 0 && col < cols && getSum(row)+getSum(col) <= k && !visited[row*cols+col] {
		return true
	}
	return false
}

// 求number数位和
// number=12 return 1+2=3
func getSum(number int) int {
	sum := 0
	for number > 0 {
		sum += number % 10
		number /= 10
	}

	return sum
}
