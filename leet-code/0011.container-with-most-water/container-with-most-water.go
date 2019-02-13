package problem0011

/*
Topic:
给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai)。
在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
说明：你不能倾斜容器，且 n 的值至少为 2。

Example:
输入: [1,8,6,2,5,4,8,3,7]
输出: 49

Thought:
穷举法是O(n^2)的复杂度，会触发leetcode的时间限制。
O(n)的复杂度的解法是，保持两个指针i,j；分别指向长度数组的首尾。如果ai 小于aj，则移动i向后（i++）。
反之，移动j向前（j--）。如果当前的area大于了所记录的area，替换之。这个想法的基础是，如果i的长度小于j，无论如何移动j，短板在i，
不可能找到比当前记录的area更大的值了，只能通过移动i来找到新的可能的更大面积。
*/
func maxArea(height []int) int {
	// 从两端开始寻找，至少保证了宽度是最大值
	i, j := 0, len(height)-1
	max := 0

	for i < j {
		a, b := height[i], height[j]
		h := min(a, b)

		area := h * (j - i)
		if max < area {
			max = area
		}

		// 朝着area具有变大的可能性方向变化。
		if a < b {
			i++
		} else {
			j--
		}
	}

	return max
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}
