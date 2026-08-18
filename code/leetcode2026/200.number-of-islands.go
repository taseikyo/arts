/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-18 15:48:32
 * @link    github.com/taseikyo
 */

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	res := 0

	var dfs func(int, int)
	dfs = func(i, j int) {
		// 边界检查：超出网格范围，或者当前格子不是陆地 '1'
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '1' {
			return
		}
		// 将当前陆地标记为已访问，避免重复计算
		grid[i][j] = '0' // 或 '2'

		// 向四个方向递归探索
		dfs(i-1, j) // 上
		dfs(i+1, j) // 下
		dfs(i, j-1) // 左
		dfs(i, j+1) // 右
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++     // 发现一个新岛屿
				dfs(i, j) // 用 DFS 淹没整个岛屿
			}
		}
	}
	return res
}
