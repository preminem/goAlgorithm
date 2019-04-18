package problem47

/*
题目：
Given a collection of numbers that might contain duplicates, return all possible unique permutations.

For example, [1,1,2] have the following unique permutations:

[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]

解题思路
本题和46. Permutations的区别在于，存在重复的数字。根据46题思路，同一个cur下，已经使用过的数字，就不能再使用了。即可
*/

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)

	n := len(nums)
	// vector 是一组可能的解答
	vector := make([]int, n)
	// taken[i] == true 表示 nums[i] 已经出现在了 vector 中
	taken := make([]bool, n)

	var ans [][]int

	makePermutation(0, n, nums, vector, taken, &ans)

	return ans
}

func makePermutation(cur, n int, nums, vector []int, taken []bool, ans *[][]int) {
	if cur == n {
		tmp := make([]int, n)
		copy(tmp, vector)
		*ans = append(*ans, tmp)
		return
	}

	used := make(map[int]bool, n-cur)

	for i := 0; i < n; i++ {

		if !taken[i] && !used[nums[i]] {
			used[nums[i]] = true

			// 准备使用 nums[i]，所以，taken[i] == true
			taken[i] = true
			// NOTICE: 是 vector[cur]
			vector[cur] = nums[i]

			makePermutation(cur+1, n, nums, vector, taken, ans)

			// 下一个循环中
			// vector[cur] = nums[i+1]
			// 所以，在这个循环中，恢复 nums[i] 自由
			taken[i] = false
		}
	}
}
