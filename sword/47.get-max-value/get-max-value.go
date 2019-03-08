package problem47

/*
题：
在一个 m*n 的棋盘中的每一个格都放一个礼物，每个礼物都有一定的价值（价值大于0）.你可以从棋盘的左上角开始拿各种里的礼物，并每次向左或者向下移动一格
，直到到达棋盘的右下角。给定一个棋盘及上面个的礼物，请计算你最多能拿走多少价值的礼物？

分析：
首先使用递归的思路进行分析，当求解到达右下角时礼物的最大总价值时，可以通过如下的递推关系进行计算
f(i,j)=ma(f(,j),f(i,j−1))+g(i,j)
其中 f(i,j)是要求解的最大值，f(i,j−1)到达(i,j)点左边点时得到最大礼物价值，而f(i−1,j)是到达(i,j)点上边点时得到最大礼物价值，g(i,j)是(i,j)点礼物的价值。
同过关系式可以使用递归的方法进行求解，但是在使用递归求解的过程中会重复求解许多的值，所以这个时候就应该使用动态规划的方式进行求解。
也就是说分析的过程如上，是从上到下递归地分析；而求解过程是从下到上循环地求解。

方法一：在动态规划求解这个问题的时候，我们找出到达每一行中每个位置的最大值。
将每一步求解出的结果都保存在一个矩阵中。那么在这个问题中就要有一个和原始矩阵等大的矩阵进行存储

方法二：实际上只需要一个与列数相同维数的一维数组就够了。
*/

func getMaxValue(matrix [][]int) int {
	if nil == matrix {
		return 0
	}

	rows := len(matrix)
	columns := len(matrix[0])
	dp := make([][]int, rows)

	for i := 0; i < rows; i++ {
		dp[i] = make([]int, columns)
		for j := 0; j < columns; j++ {
			if 0 == i && 0 == j {
				dp[i][j] = matrix[i][j]
			}
			if 0 == i && 0 != j {
				dp[i][j] = dp[i][j-1] + matrix[i][j]
			}
			if 0 != i && 0 == j {
				dp[i][j] = dp[i-1][j] + matrix[i][j]
			}
			if 0 != i && 0 != j {
				if dp[i-1][j] >= dp[i][j-1] {
					dp[i][j] = dp[i-1][j] + matrix[i][j]
				} else {
					dp[i][j] = dp[i][j-1] + matrix[i][j]
				}
			}
		}
	}
	return dp[rows-1][columns-1]
}

func getMaxValue2(matrix [][]int) int {
	if nil == matrix {
		return 0
	}

	rows := len(matrix)
	columns := len(matrix[0])
	dp := make([]int, columns)
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			left, up := 0, 0
			if 0 < i {
				up = dp[j]
			}
			if 0 < j {
				left = dp[j-1]
			}

			if left >= up {
				dp[j] = left + matrix[i][j]
			} else {
				dp[j] = up + matrix[i][j]
			}
		}
	}
	return dp[columns-1]
}