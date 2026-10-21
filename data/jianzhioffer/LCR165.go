/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-27 14:41:01
 * @link    github.com/taseikyo
 */

func crackNumber(ciphertext int) int {
	str := strconv.Itoa(ciphertext)
	n := len(str)
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		// 取出前一个字符和当前字符
		prev := str[i-2]
		cur := str[i-1]

		// 检查它们组成的两位数是否在 [10, 25] 范围内
		if prev == '1' || (prev == '2' && cur <= '5') {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[n]
}
