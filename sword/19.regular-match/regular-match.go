package problem19

/*
题:
正则表达式匹配
请实现一个函数用来匹配包好'.','*'的正则表达式。模式中'.'表示任意一个字符,'*'表示他前面的字符可以出现任意次（包含0次）
在本题中，匹配是指字符串的所有字符匹配整个模式，例如，字符串'aaa'与'a.a'和'abaca'匹配，但与'aa.a'和'ab*a'不匹配。
 */

func Math(str, pattern string) bool {
	pStr := 0
	pPattern := 0
	return matchCore(str, pattern, pStr, pPattern, len(str), len(pattern))
}

func matchCore(str, pattern string, pStr, pPattern, strLen, patternLen int) bool {
	if strLen == pStr && pPattern == patternLen{
		return true
	}

	if pStr != strLen && pPattern > patternLen {
		return false
	}

	// 防止array越界，单独判断取值
	var patternChar, strChar string
	if pPattern >= patternLen {
		patternChar = ""
	} else {
		patternChar = string(pattern[pPattern])
	}
	if pStr >= strLen {
		strChar = ""
	} else {
		strChar = string(str[pStr])
	}

	//出现*
	if (pPattern + 1) < patternLen && string(pattern[pPattern+1]) == "*" {
		//第一个字符匹配时有多种匹配方式
		if patternChar == strChar || (patternChar == "." && strLen != pStr) {
			// 进入下一个状态
			enterNext := matchCore(str, pattern, pStr+1, pPattern+2, strLen, patternLen)
			// 留在当前状态，继续匹配*
			stay := matchCore(str, pattern, pStr+1, pPattern, strLen, patternLen)
			// 略过一个*号
			ignore := matchCore(str, pattern, pStr, pPattern+2, strLen, patternLen)

			return enterNext || stay || ignore
		//第一个字符不匹配
		} else {
			// 略过一个*号
			return matchCore(str, pattern, pStr, pPattern+2, strLen, patternLen)
		}
	}

	//没有*
	if patternChar == strChar || (patternChar == "." && strLen != pStr) {
		return matchCore(str, pattern, pStr+1, pPattern+1, strLen, patternLen)
	}

	return false
}
