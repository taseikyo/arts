/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-14 16:25:35
 * @link    github.com/taseikyo
 */

func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix[0])
	for _, row := range matrix {
		if target == row[0] || target == row[n-1] {
			return true
		}
		if target < row[0] {
			return false
		}
		if target > row[n-1] {
			continue
		}
		i, j, mid := 0, n-1, 0
		for i <= j {
			mid = (j-i)/2 + i
			if row[mid] > target {
				j = mid - 1
			} else if row[mid] < target {
				i = mid + 1
			} else {
				return true
			}
		}

	}

	return false
}

func searchMatrix(matrix [][]int, target int) bool {
    // 从右上角模仿二叉搜索树的形式往左下角寻找
    for row, col := 0, len(matrix[0]) - 1; row < len(matrix) && col >= 0; {
        if matrix[row][col] > target {
            col -= 1
        } else if matrix[row][col] < target {
            row += 1
        } else {
            return true
        }
    }
    return false
}
