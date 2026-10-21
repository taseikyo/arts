/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-21 14:35:27
 * @link    github.com/taseikyo
 */

func generate(numRows int) [][]int {
	if numRows < 2 {
		return [][]int{{1}}
	}
	dp := make([][]int, numRows)
	for idx := range dp {
		dp[idx] = make([]int, idx+1)
	}
	dp[0] = []int{1}
	dp[1] = []int{1, 1}
	for r := 2; r < numRows; r++ {
		for c := 0; c <= r; c++ {
			if c == 0 || c == r {
				dp[r][c] = 1
			} else {
				dp[r][c] = dp[r-1][c] + dp[r-1][c-1]
			}
		}
	}

	return dp
}
