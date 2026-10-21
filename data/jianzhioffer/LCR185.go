/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-27 14:51:33
 * @link    github.com/taseikyo
 */

func statisticsProbability(n int) []float64 {
	// dp[i][j] 表示 i 个骰子掷出和为 j 的方案数
	// 最大和为 6*n，最小为 n
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 6*n+1)
	}

	// 初始化：1个骰子
	for j := 1; j <= 6; j++ {
		dp[1][j] = 1
	}

	// 从第2个骰子开始递推
	for i := 2; i <= n; i++ {
		// 点数和范围：从 i 到 6*i
		for j := i; j <= 6*i; j++ {
			// 当前骰子掷出 k (1~6)
			for k := 1; k <= 6; k++ {
				if j-k >= i-1 && j-k <= 6*(i-1) {
					dp[i][j] += dp[i-1][j-k]
				}
			}
		}
	}

	// 计算总方案数 = 6^n
	total := float64(1)
	for i := 0; i < n; i++ {
		total *= 6
	}

	// 收集结果（点数和从 n 到 6*n）
	res := make([]float64, 5*n+1)
	for j := n; j <= 6*n; j++ {
		res[j-n] = float64(dp[n][j]) / total
	}
	return res
}
