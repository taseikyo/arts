/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-13 16:45:06
 * @link    github.com/taseikyo
 */

func setZeroes(matrix [][]int) {
    m, n := len(matrix), len(matrix[0])
    rows := make([]bool, m)
    cols := make([]bool, n)

    // 记录哪些行和列需要置零
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 0 {
                rows[i] = true
                cols[j] = true
            }
        }
    }

    // 根据记录置零
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if rows[i] || cols[j] {
                matrix[i][j] = 0
            }
        }
    }
}

func setZeroes(matrix [][]int) {
    m, n := len(matrix), len(matrix[0])
    row0HasZero := false
    col0HasZero := false

    // 检查第一行是否有零
    for j := 0; j < n; j++ {
        if matrix[0][j] == 0 {
            row0HasZero = true
            break
        }
    }
    // 检查第一列是否有零
    for i := 0; i < m; i++ {
        if matrix[i][0] == 0 {
            col0HasZero = true
            break
        }
    }

    // 使用第一行和第一列作为标记
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0
                matrix[0][j] = 0
            }
        }
    }

    // 根据标记置零（除了第一行和第一列）
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if matrix[i][0] == 0 || matrix[0][j] == 0 {
                matrix[i][j] = 0
            }
        }
    }

    // 处理第一行
    if row0HasZero {
        for j := 0; j < n; j++ {
            matrix[0][j] = 0
        }
    }
    // 处理第一列
    if col0HasZero {
        for i := 0; i < m; i++ {
            matrix[i][0] = 0
        }
    }
}

