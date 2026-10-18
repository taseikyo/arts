/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-18 18:10:36
 * @link    github.com/taseikyo
 */

func isValid(s string) bool {
	// 用切片模拟栈
	stack := make([]rune, 0)
	// 建立右括号到左括号的映射，方便匹配
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		// 如果是右括号
		if left, ok := pairs[char]; ok {
			// 栈空 或 栈顶不匹配，直接返回 false
			if len(stack) == 0 || stack[len(stack)-1] != left {
				return false
			}
			// 匹配成功，弹出栈顶
			stack = stack[:len(stack)-1]
		} else {
			// 如果是左括号，入栈
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}
