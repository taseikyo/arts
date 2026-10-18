/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-18 16:50:59
 * @link    github.com/taseikyo
 */

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1

	for left <= right {
		mid := left + (right-left)/2
		// 将一维索引 mid 转换为二维坐标
		midVal := matrix[mid/n][mid%n]

		if midVal == target {
			return true
		} else if midVal < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
