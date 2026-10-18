/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-20 00:19:15
 * @link    github.com/taseikyo
 */

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	// 初始化答案数组，默认值为0
	answer := make([]int, n)
	// 用切片模拟单调栈，存储下标
	stack := []int{}

	for i := 0; i < n; i++ {
		// 当前温度大于栈顶下标对应的温度时，循环处理
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			// 弹出栈顶下标
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 计算等待天数
			answer[idx] = i - idx
		}
		// 将当前下标入栈
		stack = append(stack, i)
	}

	return answer
}
