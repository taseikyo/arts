/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-21 15:30:27
 * @link    github.com/taseikyo
 */

func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	// 1. 创建 (n+1) x (m+1) 的二维 DP 表
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	// 2. 初始化边界
	for i := 0; i <= n; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= m; j++ {
		dp[0][j] = j
	}

	// 3. 填充 DP 表
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(
					dp[i-1][j-1]+1, // 替换
					dp[i-1][j]+1,   // 删除
					dp[i][j-1]+1,   // 插入
				)
			}
		}
	}

	// 4. 返回结果
	return dp[n][m]
}
