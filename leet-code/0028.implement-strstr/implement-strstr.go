package problem0028

/*
Topic:
给定一个haystack字符串和一个needle字符串，在haystack字符串中找出needle字符串出现的第一个位置(从0开始)。
如果不存在，则返回 -1

Example:
示例 1:
输入: haystack = "hello", needle = "ll"
输出: 2

示例 2:
输入: haystack = "aaaaa", needle = "bba"
输出: -1

说明:
当 needle是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
对于本题而言，当 needle 是空字符串时我们应当返回0。这与C语言的strstr()以及Java的indexOf()定义相符。
 */

func strStr(haystack string, needle string) int {
	hlen, nlen := len(haystack), len(needle)
	// 当hlen等于nlen的时候，需要i == 0
	for i := 0; i <= hlen-nlen; i++ {
		if haystack[i:i+nlen] == needle {
			return i
		}
	}

	return -1
}

func strStrMy(haystack string, needle string) int {
	if needle == ""{
		return 0
	}
	for i:=0;i<len(haystack)-len(needle)+1;i++{
		for j:=0;j<len(needle);j++{
			if haystack[i+j] != needle[j]{
				break
			}
			if j == len(needle) - 1{
				return i
			}
		}
	}
	return -1
}


