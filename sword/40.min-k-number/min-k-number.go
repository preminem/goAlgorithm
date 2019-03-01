package problem40

/*
题：
最小的k个数
输入n个整数，找出其中最小的k个数

分析：
解法一：先排序，复杂度O(nlogn),效率太慢
解法二：受快速排序启发，复杂度O(n),但需改变数组
解法三：复杂度为O（nlogk)的算法，利用堆排序，适合海量数据
 */

//方法二
//选出k个最小元素，k为1~len(s)
func SelectKMin(s []int, k int) []int {
	lo, hi := 0, len(s)-1
	for {
		j := partition(s, lo, hi)
		if j < k {
			lo = j + 1
		} else if j > k {
			hi = j - 1
		} else {
			return s[:k]
		}
	}
}

func partition(s []int, lo, hi int) int {
	i, j := lo, hi+1
	for {
		for {
			i++
			if i == hi || s[i] > s[lo] {
				break
			}
		}
		for {
			j--
			if j == lo || s[j] <= s[lo] {
				break
			}
		}
		if i >= j {
			break
		}
		swap(s, i, j)
	}
	swap(s, lo, j)
	return j
}

func swap(s []int, i int, j int) {
	s[i], s[j] = s[j], s[i]
}