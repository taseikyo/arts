/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-21 14:01:29
 * @link    github.com/taseikyo
 */

func exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])
	// 初始化访问标记矩阵
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	// 定义深度优先搜索函数
	var dfs func(r, c, idx int) bool
	dfs = func(r, c, idx int) bool {
		// 1. 边界条件与剪枝
		if r < 0 || r >= rows || c < 0 || c >= cols ||
			visited[r][c] || board[r][c] != word[idx] {
			return false
		}
		// 2. 找到完整单词
		if idx == len(word)-1 {
			return true
		}

		// 3. 标记当前格子为已访问
		visited[r][c] = true

		// 4. 向四个方向递归探索
		// 使用方向数组让代码更简洁
		directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, d := range directions {
			nr, nc := r+d[0], c+d[1]
			if dfs(nr, nc, idx+1) {
				return true
			}
		}

		// 5. 回溯：撤销当前格子的访问标记
		visited[r][c] = false
		return false
	}

	// 从每个格子开始尝试
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] == word[0] && dfs(r, c, 0) {
				return true
			}
		}
	}
	return false
}
