/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-21 00:44:28
 * @link    github.com/taseikyo
 */

func generateParenthesis(n int) []string {
	var res []string

	// 定义回溯函数
	// left: 已使用的左括号数, right: 已使用的右括号数, path: 当前构建的字符串
	var backtrack func(left, right int, path string)
	backtrack = func(left, right int, path string) {
		// 终止条件：当字符串长度达到 2*n 时，说明找到一个有效组合
		if len(path) == 2*n {
			res = append(res, path)
			return
		}

		// 1. 尝试添加左括号：只要左括号数量不超过 n
		if left < n {
			backtrack(left+1, right, path+"(") 
		}

		// 2. 尝试添加右括号：只要右括号数量少于左括号数量
		if right < left {
			backtrack(left, right+1, path+")") 
		}
	}

	// 从空字符串开始回溯
	backtrack(0, 0, "")
	return res
}
