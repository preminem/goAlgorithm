package problem0017

/*
Topic:
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

Example:
输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
说明:尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
*/

var m = map[byte][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	ret := []string{""}

	// 让digits中所有的数字都有机会进行迭代。
	for i := 0; i < len(digits); i++ {
		temp := []string{}
		// 让ret中的每个元素都能添加新的字母。
		for j := 0; j < len(ret); j++ {
			// 把digits[i]所对应的字母，多次添加到ret[j]的末尾
			for k := 0; k < len(m[digits[i]]); k++ {
				temp = append(temp, ret[j]+m[digits[i]][k])
			}
		}

		ret = temp
	}

	return ret
}
