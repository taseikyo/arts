/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-27 14:31:39
 * @link    github.com/taseikyo
 */

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	// 创建 (m+1) x (n+1) 的 DP 表
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true // 空串匹配空串

	// 初始化：处理 p 可以匹配空字符串 s 的情况
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	// 填充 DP 表
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				// 情况1: '*' 匹配零次
				dp[i][j] = dp[i][j-2]
				// 情况2: '*' 匹配一次或多次
				// 前提是 p[j-2] 能与 s[i-1] 匹配
				if p[j-2] == '.' || p[j-2] == s[i-1] {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else {
				// 普通字符或 '.' 的匹配
				if p[j-1] == '.' || p[j-1] == s[i-1] {
					dp[i][j] = dp[i-1][j-1]
				}
			}
		}
	}

	return dp[m][n]
}
