/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-23 20:44:01
 * @link    github.com/taseikyo
 */

func mechanicalAccumulator(n int) int {
	ans := 0

	// 定义一个递归函数，它执行累加并返回一个 bool 值
	// 这个 bool 值只是为了满足 && 运算符的语法要求
	var helper func(int) bool
	helper = func(num int) bool {
		// 1. 执行累加：将当前数字 num 加到结果 ans 中
		ans += num

		// 2. 递归终止与继续的关键：
		//    - 当 num > 0 时，条件为 true，会执行 helper(num-1)，继续递归
		//    - 当 num == 0 时，条件为 false，发生短路，helper(num-1) 不会执行，递归终止
		//    最终返回的 bool 值本身没有实际意义，只是为了符合语法
		return num > 0 && helper(num-1)
	}

	helper(n)
	return ans
}
