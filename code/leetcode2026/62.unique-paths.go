/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-21 15:11:54
 * @link    github.com/taseikyo
 */

func uniquePaths(m int, n int) int {
	// 1. 创建二维 DP 数组
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 2. 初始化边界：第一行和第一列都是 1
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	// 3. 填充 DP 表格
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	// 4. 返回右下角的值
	return dp[m-1][n-1]
}
