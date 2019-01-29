package problem0001

/*
Topic:
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:
Given nums = [2, 7, 11, 15], target = 9,
Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

Thought：
直接使用两层for循环复杂度为O(n^2)
a+b=target 即 b=target-a
更好的解决方法是使用map[<整数>]<整数的序号>保存已经遍历过的整数
*/
func twoSum(nums []int, target int) []int {
	// index 负责保存map[整数]整数的序列号
	index := make(map[int]int, len(nums))

	// 通过 for 循环，获取b的序列号
	for i, b := range nums {
		// 通过查询map，获取a = target - b的序列号
		if j, ok := index[target-b]; ok {
			// ok 为 true
			// 说明在i之前，存在 nums[j] == a
			return []int{j, i}
			// 注意，顺序是j，i
		}

		// 把b和i的值，存入map
		index[b] = i
	}

	return nil
}
