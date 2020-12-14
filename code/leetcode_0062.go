/**
 * @date    2020-12-09 10:53:17
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @link 	github.com/taseikyo
 */

func uniquePaths(m int, n int) int {
	board := make([][]int, m, m)
	for i := 0; i < m; i++ {
		col := make([]int, n, n)
		board[i] = col
	}
	for i := 0; i < n; i++ {
		board[0][i] = 1
	}
	for i := 0; i < m; i++ {
		board[i][0] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			board[i][j] = board[i-1][j] + board[i][j-1]
		}
	}
	return board[m-1][n-1]
}