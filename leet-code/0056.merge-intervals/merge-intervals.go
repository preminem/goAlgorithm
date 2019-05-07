package problem0056

/*
题：
给出一个区间的集合，请合并所有重叠的区间。

示例 1:
输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

示例 2:
输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。

思路：
先对 intervals 进行排序，按照 Start 递增
依次处理重叠的情况。
*/
import (
	"sort"
)

// Interval Definition for an interval.
type Interval struct {
	Start int
	End   int
}

func merge(its []Interval) []Interval {
	if len(its) <= 1 {
		return its
	}

	sort.Slice(its, func(i int, j int) bool {
		return its[i].Start < its[j].Start
	})

	res := make([]Interval, 0, len(its))

	temp := its[0]
	for i := 1; i < len(its); i++ {
		if its[i].Start <= temp.End {
			temp.End = max(temp.End, its[i].End)
		} else {
			res = append(res, temp)
			temp = its[i]
		}
	}
	res = append(res, temp)

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
