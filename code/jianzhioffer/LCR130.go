/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-27 15:36:51
 * @link    github.com/taseikyo
 */

func wardrobeFinishing(m int, n int, cnt int) int {
	// visited 记录哪些格子已经被访问过
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	// 计算数位之和
	digitSum := func(x int) int {
		sum := 0
		for x > 0 {
			sum += x % 10
			x /= 10
		}
		return sum
	}

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		// 边界检查：越界、已访问、数位和超标
		if r < 0 || r >= m || c < 0 || c >= n {
			return 0
		}
		if visited[r][c] {
			return 0
		}
		if digitSum(r)+digitSum(c) > cnt {
			return 0
		}

		// 标记当前格子为已访问
		visited[r][c] = true

		// 只向右和向下探索（从原点出发，这两个方向足够覆盖所有可达格子）
		return 1 + dfs(r+1, c) + dfs(r, c+1)
	}

	return dfs(0, 0)
}
