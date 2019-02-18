package problem0020

/*
Topic:
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串

Example:
Example 1:
Input: "()"
Output: true

Example 2:
Input: "()[]{}"
Output: true

Example 3:
Input: "(]"
Output: false

Example 4:
Input: "([)]"
Output: false

Example 5:
Input: "{[]}"
Output: true

Thought:
1.初始化栈 S。
2.一次处理表达式的每个括号。
3.如果遇到开括号，我们只需将其推到栈上即可。这意味着我们将稍后处理它，让我们简单地转到前面的子表达式。
4.如果我们遇到一个闭括号，那么我们检查栈顶的元素。如果栈顶的元素是一个相同类型的左括号，那么我们将它从栈中弹出并继续处理。否则，这意味着表达式无效。
5.如果到最后我们剩下的栈中仍然有元素，那么这意味着表达式无效。
*/

func isValid(s string) bool {
	size := len(s)

	stack := make([]byte, size)
	top := 0

	for i := 0; i < size; i++ {
		c := s[i]
		switch c {
		case '(':
			stack[top] = c + 1 // '('+1 is ')'
			top++
		case '[', '{':
			stack[top] = c + 2 // '{'+2 is '}'    '['+2 is ']'
			top++
		case ')', ']', '}':
			if top > 0 && stack[top-1] == c {
				top--
			} else {
				return false
			}
		}
	}

	return top == 0
}
