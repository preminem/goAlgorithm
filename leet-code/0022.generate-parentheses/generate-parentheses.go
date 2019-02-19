package problem0022

/*
Topic:
给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。

Example:
given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]

Thought:
在当前局面下，你有若干种选择。那么尝试每一种选择。如果已经发现某种选择肯定不行（因为违反了某些限定条件），就返回；如果某种选择试到最后发现是正确解，就将其加入解集

需要从两个方面去思考：1. 选择与限制；2.结束条件

对于这道题，在任何时刻，你都有两种选择：
加左括号。
加右括号。

同时有以下限制：
如果左括号已经用完了，则不能再加左括号了。
如果已经出现的右括号和左括号一样多，则不能再加右括号了。因为那样的话新加入的右括号一定无法匹配。

结束条件是：左右括号都已经用完。

结束后的正确性：左右括号用完以后，一定是正确解。因为1. 左右括号一样多，2. 每个右括号都一定有与之配对的左括号。因此一旦结束就可以加入解集。
递归函数传入参数：限制和结束条件中有“用完”和“一样多”字样，因此你需要知道左右括号的数目。

当然你还需要知道当前局面substr和解集res。

因此，把上面的思路拼起来就是代码：

if (左右括号都已用完) {
  加入解集，返回
}
//否则开始试各种选择
if (还有左括号可以用) {
  加一个左括号，继续递归
}
if (还有右括号可以用，且，右括号小于左括号) {
  加一个右括号，继续递归
}
*/

func generateParenthesis(n int) []string {
	res := make([]string, 0, n*n)
	bytes := make([]byte, n*2)
	dfs(n, n, 0, bytes, &res)
	return res
}

func dfs(left, right, idx int, bytes []byte, res *[]string) {
	// 所有符号都添加完毕
	if left == 0 && right == 0 {
		*res = append(*res, string(bytes))
		return
	}

	// "(" 不用担心匹配问题，
	// 只要 left > 0 就可以直接添加
	if left > 0 {
		bytes[idx] = '('
		dfs(left-1, right, idx+1, bytes, res)
	}

	// 想要添加 ")" 时
	// 需要 left < right，
	// 即在 bytes[:idx] 至少有一个 "(" 可以与 这个 ")" 匹配
	if right > 0 && left < right {
		bytes[idx] = ')'
		dfs(left, right-1, idx+1, bytes, res)
	}
}

//func generateParenthesis(n int) []string {
//	res := new([]string)
//	backtrack("",res,n,0,0)
//	return *res
//}
//
//func backtrack(s string, res *[]string,n int, left int, right int){
//	if len(s) == 2*n{
//		*res = append(*res,s)
//		return
//	}
//	if left < n{
//		backtrack(s+"(",res,n,left+1,right)
//	}
//	if right < left{
//		backtrack(s+")",res,n,left,right+1)
//	}
//}
