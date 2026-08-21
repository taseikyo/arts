/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-21 14:18:00
 * @link    github.com/taseikyo
 */

func solveNQueens(n int) [][]string {
    var result [][]string
    // 用三个布尔数组分别记录被占用的列、主对角线（r-c）、副对角线（r+c）
    // 主对角线数量为 2*n-1，副对角线数量为 2*n-1
    cols := make([]bool, n)
    diag1 := make([]bool, 2*n-1) // 主对角线，索引为 r-c + (n-1)
    diag2 := make([]bool, 2*n-1) // 副对角线，索引为 r+c

    // queens 数组记录每一行皇后所在的列，queens[r] = c
    queens := make([]int, n)

    var backtrack func(row int)
    backtrack = func(row int) {
        // 如果已经成功放置了 n 个皇后，说明找到了一个解
        if row == n {
            // 根据 queens 数组构建棋盘字符串
            board := make([]string, n)
            for r := 0; r < n; r++ {
                rowStr := make([]byte, n)
                for c := 0; c < n; c++ {
                    if queens[r] == c {
                        rowStr[c] = 'Q'
                    } else {
                        rowStr[c] = '.'
                    }
                }
                board[r] = string(rowStr)
            }
            result = append(result, board)
            return
        }

        // 尝试在当前 row 的每一列放置皇后
        for col := 0; col < n; col++ {
            // 计算当前格子所在的两条对角线的索引
            d1Idx := row - col + n - 1 // 主对角线
            d2Idx := row + col         // 副对角线

            // 如果列、两条对角线都没有被占用，则可以选择
            if !cols[col] && !diag1[d1Idx] && !diag2[d2Idx] {
                // 做选择
                queens[row] = col
                cols[col] = true
                diag1[d1Idx] = true
                diag2[d2Idx] = true

                // 递归进入下一行
                backtrack(row + 1)

                // 撤销选择（回溯）
                cols[col] = false
                diag1[d1Idx] = false
                diag2[d2Idx] = false
            }
        }
    }

    backtrack(0)
    return result
}
