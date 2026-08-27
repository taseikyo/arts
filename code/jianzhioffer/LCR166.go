/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-27 14:51:05
 * @link    github.com/taseikyo
 */

func jewelleryValue(frame [][]int) int {
	m, n := len(frame), len(frame[0])
	for r := 1; r < m; r++ {
		frame[r][0] += frame[r-1][0]
	}
	for c := 1; c < n; c++ {
		frame[0][c] += frame[0][c-1]
	}
	for r := 1; r < m; r++ {
		for c := 1; c < n; c++ {
			frame[r][c] += max(frame[r-1][c], frame[r][c-1])
		}
	}

	return frame[m-1][n-1]
}
