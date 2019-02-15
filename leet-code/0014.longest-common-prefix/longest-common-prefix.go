package problem0014

/*
Topic:
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。

Thought:
方法一：水平扫描，LCP(S1…Sn)=LCP(LCP(LCP(S1,S2),S3),…Sn)
方法二：从前往后枚举字符串的每一列，先比较每个字符串相同列上的字符（即不同字符串相同下标的字符）然后再进行对下一列的比较。
方法三：分治，LCP(S1…Sn)=LCP(LCP(S1​…Sk),LCP(Sk+1…Sn))
方法四：二分查找法，这个想法是应用二分查找法找到所有字符串的公共前缀的最大长度 L。算法的查找区间是(0...minLen)(0…minLen)，
其中minLen是输入数据中最短的字符串的长度，同时也是答案的最长可能长度。每一次将查找区间一分为二，然后丢弃一定不包含最终答案的那一个。
*/

//找到最短的字符串再逐列比较
func longestCommonPrefix(strs []string) string {
	short := shortest(strs)

	for i, r := range short {
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != byte(r) {
				return strs[j][:i]
			}
		}
	}

	return short
}

func shortest(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	res := strs[0]
	for _, s := range strs {
		if len(res) > len(s) {
			res = s
		}
	}

	return res
}