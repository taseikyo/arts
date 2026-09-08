/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-09-06 23:40:06
 * @link    github.com/taseikyo
 */

func numDistinct(s string, t string) int {
	// 获取两个字符串的长度
	m, n := len(s), len(t)

	// 如果 t 的长度大于 s，那一定不可能构成，直接返回 0
	// 但 DP 表也能处理这种情况，不过为了效率，可以提前判断[reference:10]
	if n > m {
		return 0
	}

	// 1. 创建 DP 表，维度为 (n+1) x (m+1)
	// dp[i][j] 表示 t 的前 i 个字符在 s 的前 j 个字符中出现的次数[reference:11]
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	// 2. 初始化：当 i == 0 时，即 t 为空字符串，方案数为 1[reference:12]
	for j := 0; j <= m; j++ {
		dp[0][j] = 1
	}

	// 3. 填充 DP 表
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if t[i-1] == s[j-1] {
				// 情况二：字符匹配，可以选择“使用”或“跳过”[reference:13]
				dp[i][j] = dp[i][j-1] + dp[i-1][j-1]
			} else {
				// 情况一：字符不匹配，只能“跳过”[reference:14]
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	// 4. 最终结果：t 的全部 n 个字符在 s 的全部 m 个字符中出现的次数
	return dp[n][m]
}
