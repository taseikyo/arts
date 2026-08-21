/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-21 14:25:11
 * @link    github.com/taseikyo
 */

func climbStairs(n int) int {
	if n < 2 {
		return 1
	}
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n-1]
}
