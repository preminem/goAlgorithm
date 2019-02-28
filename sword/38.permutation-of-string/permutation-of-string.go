package problem38

/*
题：
字符串的全排列
输入一个字符串，打印这个字符串所有字符的全排列 例如，输入abc,则打印a,b,c所能排列出来的所有字符串abc,acb,bac,bca,cab,cba。

分析：
分成两步看
1.求出所有可能在第一个位置的字符
2.固定这个字符求后面字符的排列
 */
import "fmt"


func Permutation(str string) {
	if str == "" {
		return
	}

	permutation([]rune(str), 0)
}
//sIndex指向当前我们执行排列操作的字符串的第一个字符。在每一次递归的时候，我们从sIndex向后扫描每一个字符（指针i指向的字符）。
//在交换sIndex和i指向的字符后，我们再对sIdex之后的字符串递归地进行排列操作。
func permutation(str []rune, sIndex int) {
	if sIndex == len(str) {
		fmt.Printf("%s\n", string(str))
	} else {
		for i:= sIndex; i != len(str); i++ {
			str[sIndex], str[i] = str[i], str[sIndex]
			permutation(str, sIndex + 1)
			str[sIndex], str[i] = str[i], str[sIndex]
		}
	}
}