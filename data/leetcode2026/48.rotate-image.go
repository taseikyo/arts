/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-14 15:56:49
 * @link    github.com/taseikyo
 */

func rotate(matrix [][]int) {
	// 右上 - 左下
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 水平对调
	for _, row := range matrix {
		i, j := 0, len(row) - 1
		for i < j {
			row[i], row[j] = row[j], row[i]
			i++
			j--
		}
	}
}
