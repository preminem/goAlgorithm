package problem49

/*
题目:
字母异位词分组
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

For example, given: ["eat", "tea", "tan", "ate", "nat", "bat"],
Return:

[
  ["ate", "eat","tea"],
  ["nat","tan"],
  ["bat"]
]

说明：
所有输入均为小写字母。
不考虑答案输出的顺序

思路一：
当且仅当它们的排序字符串相等时，两个字符串是字母异位词。
维护一个映射，其中每个键K是一个排序字符串，每个值是初始输入的字符串列表，排序后等于K。

思路二：
当且仅当它们的字符计数（每个字符的出现次数）相同时，两个字符串是字母异位词。
我们可以将每个字符串s转换为字符数count，由26个非负整数组成，表示a，b，c的数量等。我们使用这些计数作为哈希映射的基础。
（可以用质数表示26个字母，把字符串的各个字母相乘，这样可保证字母异位词的乘积必定是相等的。）

*/

func groupAnagrams(ss []string) [][]string {
	tmp := make(map[int][]string, len(ss)/2)
	for _, s := range ss {
		c := encode(s)
		tmp[c] = append(tmp[c], s)
	}

	res := make([][]string, 0, len(tmp))
	for _, v := range tmp {
		res = append(res, v)
	}

	return res
}

// prime 与 A～Z 对应，英文中出现概率越大的字母，选用的质数越小
var prime = []int{5, 71, 37, 29, 2, 53, 59, 19, 11, 83, 79, 31, 43, 13, 7, 67, 97, 23, 17, 3, 41, 73, 47, 89, 61, 101}

func encode(s string) int {
	res := 1
	for i := range s {
		res *= prime[s[i]-'a']
	}
	return res
}
