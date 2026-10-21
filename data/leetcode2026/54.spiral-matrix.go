/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-13 17:25:06
 * @link    github.com/taseikyo
 */

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	rowLen, colLen := len(matrix), len(matrix[0])
	row, col := 0, 0
	res := make([]int, 0, rowLen*colLen)

	for rowLen > 0 && colLen > 0 {
		// 右
		for i := col; i < col+colLen; i++ {
			res = append(res, matrix[row][i])
		}
		// 下
		for i := row + 1; i < row+rowLen; i++ {
			res = append(res, matrix[i][col+colLen-1])
		}
		// 左：确保不止一行
		if rowLen > 1 {
			for i := col + colLen - 2; i >= col; i-- {
				res = append(res, matrix[row+rowLen-1][i])
			}
		}
		// 上：确保不止一列
		if colLen > 1 {
			for i := row + rowLen - 2; i > row; i-- {
				res = append(res, matrix[i][col])
			}
		}

		row++
		col++
		rowLen -= 2
		colLen -= 2
	}
	return res
}
