package problem0003

/*
Topic:
Given a string, find the length of the longest substring without repeating characters.

Example:
Given "abcabcbb", the answer is "abc", which the length is 3.
Given "bbbbb", the answer is "b", with the length of 1.
Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

Thought：
我们使用 HashSet[<字符>]<序号> 来判断窗口中的字符上一次出现的位置是否在窗口之中。
然后我们向右侧滑动索引i，如果它不在 HashSet 中，我们会继续滑动i。直到 s[i] 已经存在于 HashSet 中。
此时，我们找到的没有重复字符的最长子字符串将会以索引left开头。如果我们对所有的left这样做，就可以得到答案。
我们可以用一个整数数组作为直接访问表来替换 Map,常用的表如下所示：
int [26] 用于字母 ‘a’ - ‘z’或 ‘A’ - ‘Z’
int [128] 用于ASCII码
int [256] 用于扩展ASCII码

Complexity:
时间复杂度：O(n)，索引i将会迭代n次。
空间复杂度：O(m)，m是字符集的大小。

Summary:
需要一个窗口向右滑动搜索，需要一个映射保存已经出现过的字符位置i避免查询，如果i在窗口内则窗口左边右移至i+1。
*/

func lengthOfLongestSubstring(s string) int {
	// location[s[i]] == j 表示：
	// s中第i个字符串，上次出现在s的j位置，所以，在s[j+1:i]中没有s[i]
	// location[s[i]] == -1 表示： s[i] 在s中第一次出现
	location := [256]int{} // 只有256长是因为，假定输入的字符串只有ASCII字符
	for i := range location {
		location[i] = -1 // 先设置所有的字符都没有见过
	}

	maxLen, left := 0, 0

	for i := 0; i < len(s); i++ {
		// 说明s[i]已经在s[left:i+1]中重复了
		// 并且s[i]上次出现的位置在location[s[i]]
		if location[s[i]] >= left {
			left = location[s[i]] + 1 // 在s[left:i+1]中去除s[i]字符及其之前的部分
		} else if i+1-left > maxLen {
			// fmt.Println(s[left:i+1])
			maxLen = i + 1 - left
		}
		location[s[i]] = i
	}

	return maxLen
}
