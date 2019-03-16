package problem60

import (
	"fmt"
	"math"
)

/*
题目：n个骰子的点数
把n个骰子仍在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值的出现概率。

思路标签：
算法：递归、动态规划

解答：
1. 基于递归的方法，效率低
我们将大问题n个骰子的点数和，转化为求第1个和后面n-1个骰子；
同时用一个数组来保存所有可能和出现的次数，注意是和出现的次数，即每次该和出现一次，对应的位置上就+1；
利用递归去求解，每一个骰子都有6种可能性，排列组合总共有6^n种可能；
结束条件是每次得到最后一个骰子的结果；
最后求概率注意int和double的转换。
*/

var maxValue = 6

func PrintProbabilityRecursive(number int) {
	if number < 1 {
		return
	}
	maxSum := number * maxValue
	probabilities := make([]int, maxSum-number+1)
	probability(number, probabilities)
	total := math.Pow(float64(maxValue), float64(number))
	for i := number; i <= maxSum; i++ {
		ratio := float64(probabilities[i-number]) / total
		fmt.Printf("%d:%e\n", i, ratio)
	}
}

func probability(number int, probabilities []int) {
	for i := 1; i <= maxValue; i++ {
		probabilityCore(number, number, i, probabilities)
	}
}

//original:一个几个筛子，current:目前处理到倒数第几个筛子，sum：前面的筛子累计和是多少
func probabilityCore(original, current, sum int, probabilities []int) {
	//当处理到倒数第一个筛子的时候，对应点数的个数+1
	if current == 1 {
		probabilities[sum-original] += 1
	} else {
		for i := 1; i <= maxValue; i++ {
			probabilityCore(original, current-1, i+sum, probabilities)
		}
	}
}

/*
2. 基于循环求骰子点数，性能较好（动态规划）
使用动态规划的思想其时间复杂度总是要小于递归的方法；
我们设立数组，用数组中的第n个数表示骰子点数和为n的次数；
第k次投掷骰子的数可能为1~6中的任意一个数，如果我们假设第k次投掷骰子最终所有的和为n，那么和为n的次数就为前一次投掷（第k-1次投掷）和为n-1、n-2、n-3、n-4、n-5、n-6的次数的总和。
同时知道第1次投掷和为1,2,3,4,5,6的次数均为1；同时第k次投掷时，和为0、1、2…k-1将不会存在；
我们使用两个数组来交替进行，一个用来保存上一次投掷的和的次数，另一个以对方为基础用来计算当前投掷和的次数。每次用flag来交替。
*/

func PrintProbability(number int) {
	if number < 1 {
		return
	}
	probabilities := [2][]int{make([]int, maxValue*number+1), make([]int, maxValue*number+1)}
	flag := 0
	//填充第一轮的次数
	for i := 1; i <= maxValue; i++ {
		probabilities[flag][i] = 1
	}
	//从第二轮开始
	for k := 2; k <= number; k++ {
		//第k次投掷时和为0、1、2…k-1将不会存在
		for i := 0; i < k; i++ {
			probabilities[1-flag][i] = 0
		}
		//第k次投掷和最小为k，最大为6*k
		//和为n的次数为前一次投掷和为n-1、n-2、n-3、n-4、n-5、n-6的次数的总和
		for i := k; i <= maxValue*k; i++ {
			probabilities[1-flag][i] = 0
			for j := 1; j <= i && j <= maxValue; j++ {
				probabilities[1-flag][i] += probabilities[flag][i-j]
			}
		}
		flag = 1 - flag
	}
	total := math.Pow(float64(maxValue), float64(number))
	maxSum := number * maxValue
	for i := number; i <= maxSum; i++ {
		ratio := float64(probabilities[flag][i]) / total
		fmt.Printf("%d:%e\n", i, ratio)
	}
}
