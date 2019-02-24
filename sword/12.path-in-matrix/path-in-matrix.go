package problem12

/*
题：
矩阵中的路径
请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符路径。
路径可以从矩阵任意一格开始，每一步可以在矩阵中向左、右、上下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。
例如，在下面的3*4的矩阵中包含一条字符串"bfce"的路径。但矩阵中不包含"abfb"的路径。

a b t g
c f c s
j d e h

分析：
回溯法
*/

func HasPath(matrix string, rows, cols int, str string) bool {
	// 非法输入返回false
	if matrix == "" || rows < 1 || cols < 1 || str == "" {
		return false
	}

	//用来表示格子是否已经进入了路径
	visited := make([]bool, rows*cols)

	pathLength := 0

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// 起始点
			if hasPathCore(matrix, rows, cols, row, col, &pathLength, str, visited) {
				return true
			}
		}
	}
	return false
}

func hasPathCore(matrix string, rows, cols, row, col int, pathLength *int, str string, visited []bool) bool {
	hasPath := false

	// 长度等于查找字符长度时即推出
	if *pathLength == len(str) {
		return true
	}

	if row >= 0 && row < rows && col >= 0 && col < cols && matrix[row*cols+col] == str[*pathLength] && !visited[row*cols+col] {
		*pathLength++
		visited[row*cols+col] = true

		// 向上下左右前进一步，匹配成功则为true
		hasPath = hasPathCore(matrix, rows, cols, row, col-1, pathLength, str, visited) || hasPathCore(matrix, rows, cols, row-1, col, pathLength, str, visited) || hasPathCore(matrix, rows, cols, row, col+1, pathLength, str, visited) || hasPathCore(matrix, rows, cols, row+1, col, pathLength, str, visited)

		// 向前一步没有匹配的路径，失败则退回
		if !hasPath {
			*pathLength--
			visited[row*cols+col] = false
		}
	}

	return hasPath
}
