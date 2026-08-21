/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-21 14:58:07
 * @link    github.com/taseikyo
 */

func numSquares(n int) int {
	// 初始化 dp 数组，长度为 n+1
	// dp[i] 表示组成数字 i 所需的最少完全平方数个数
	dp := make([]int, n+1)

	// 从 1 开始遍历到 n
	for i := 1; i <= n; i++ {
		// 最坏情况：全由 1 组成，需要 i 个
		minCount := i
		// 尝试减去所有可能的完全平方数 j*j
		for j := 1; j*j <= i; j++ {
			// 状态转移
			if dp[i-j*j]+1 < minCount {
				minCount = dp[i-j*j] + 1
			}
		}
		dp[i] = minCount
	}

	return dp[n]
}

import "math"

func numSquares(n int) int {
	// 1. 判断是否为 1
	sqrt := int(math.Sqrt(float64(n)))
	if sqrt*sqrt == n {
		return 1
	}

	// 2. 判断是否为 4
	temp := n
	for temp%4 == 0 {
		temp /= 4
	}
	if temp%8 == 7 {
		return 4
	}

	// 3. 判断是否为 2
	for i := 1; i*i <= n; i++ {
		j := int(math.Sqrt(float64(n - i*i)))
		if i*i+j*j == n {
			return 2
		}
	}

	// 4. 否则为 3
	return 3
}
