package problem57

/*
题目一：和为s的两个数字
输入一个递增排序的数组和一个数字，在数组中查找两个数，使得他们的和正好是s，如果有多对数字的和等于s，则输出任意一对即可。

分析：
需要注意题目中给出的条件：递增排序数组。要利用数组排序的性质；
1.使用两个指针，一个指针指向数组的第一个元素，一个数组指向最后一个元素；
2.如果二者的和小于目标和，小指针向前移动；如果二者的和大于目标和，则大指针向后移动。
算法是基于递增的性质，如果不是递增则使用hash表来解决。
 */

import "fmt"

func findNumbers(data []int, sum int) (int, int, bool) {
	var found bool
	if nil == data || 1 > len(data) {
		return 0, 0, found
	}

	ahead, behind := 0, len(data)-1
	for ahead < behind {
		currentSum := data[ahead] + data[behind]
		if currentSum == sum {
			found = true
			return data[ahead], data[behind], found
		} else if currentSum < sum {
			ahead++
		} else {
			behind--
		}
	}
	return 0, 0, found
}

/*
题目二：和为s的连续正数序列
输入一个正数s，输出所有和为s的连续正数序列（至少含有两个数）。

分析：
还是以上面类似的思路，使用双指针来解决；
小指针初始为1，大指针初始为2；
如果从小到大的数字的和大于目标和，则抛弃小指针当前值，也就是小指针加1；如果从小到大数字的和小于目标和，则增加大指针；
注意循环的边界条件：samll
 */
func findContinuousSequence(sum int) {
	if 3 > sum {
		return
	}

	middle := (sum + 1) / 2
	small, big := 1, 2
	currentSum := small + big
	for small < middle {
		if currentSum == sum {
			printContinuousSequence(small, big)
		}

		for currentSum > sum && small < middle {
			currentSum -= small
			small++

			if currentSum == sum {
				printContinuousSequence(small, big)
			}
		}
		big++
		currentSum += big
	}
}

func printContinuousSequence(small, big int) {
	for i := small; i <= big; i++ {
		fmt.Printf("%v ", i)
	}
	fmt.Println()
}