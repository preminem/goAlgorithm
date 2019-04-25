package problem0046

/*
题目：
Given a collection of distinct integers, return all possible permutations.

Example:

Input: [1,2,3]
Output:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
思路：
递归函数用于求得每一位的解，for循环用来遍历输入的每一位
*/

func permute(nums []int) [][]int {
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

	for i := 0; i < n; i++ {
		if !taken[i] {
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
