/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-31 15:59:26
 * @link    github.com/taseikyo
 */

func removeKdigits(num string, k int) string {
	// 如果要删除的数量大于等于数字长度，结果为空，返回 "0"
	if k >= len(num) {
		return "0"
	}

	// 用切片模拟单调栈
	stack := make([]byte, 0, len(num))

	for i := 0; i < len(num); i++ {
		c := num[i]
		// 当还可以删除 (k>0) 且栈不为空，且栈顶元素大于当前元素时
		// 弹出栈顶元素，相当于删除这个较大的数字
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > c {
			stack = stack[:len(stack)-1]
			k--
		}
		// 处理前导 0: 如果栈为空且当前字符是 '0'，则跳过，不进行入栈
		if c == '0' && len(stack) == 0 {
			continue
		}
		stack = append(stack, c)
	}

	// 如果遍历完数字后，k 仍然大于 0（例如 num 是 "12345"）
	// 则需要从栈顶（末尾）继续删除元素，直到 k 为 0
	for k > 0 && len(stack) > 0 {
		stack = stack[:len(stack)-1]
		k--
	}

	// 如果最终栈为空，返回 "0"；否则将栈转为字符串返回
	if len(stack) == 0 {
		return "0"
	}
	return string(stack)
}
