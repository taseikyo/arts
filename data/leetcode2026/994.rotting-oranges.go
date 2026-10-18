/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-18 16:13:49
 * @link    github.com/taseikyo
 */

func orangesRotting(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	// 队列，用于BFS
	queue := make([][2]int, 0)
	freshCount := 0

	// 1. 初始化：统计新鲜橘子，并将所有腐烂橘子入队
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				freshCount++
			} else if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	// 如果没有新鲜橘子，直接返回0
	if freshCount == 0 {
		return 0
	}

	// 四个方向的移动向量
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	minutes := 0

	// 2. 多源BFS
	for len(queue) > 0 && freshCount > 0 {
		// 关键：记录当前队列的长度，这代表“当前这一分钟”要处理的腐烂橘子
		levelSize := len(queue)
		// 在这一分钟内，只处理 levelSize 个橘子
		for i := 0; i < levelSize; i++ {
			// 取出队首元素
			cur := queue[0]
			queue = queue[1:]
			x, y := cur[0], cur[1]

			// 尝试向四个方向扩散
			for _, d := range directions {
				nx, ny := x+d[0], y+d[1]
				// 如果相邻格子是新鲜橘子
				if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == 1 {
					// 让它腐烂
					grid[nx][ny] = 2
					freshCount--
					// 将这个新腐烂的橘子加入队列，作为下一分钟的“源头”
					queue = append(queue, [2]int{nx, ny})
				}
			}
		}
		// 处理完当前这一分钟的所有橘子后，时间+1
		minutes++
	}

	// 3. 如果还有新鲜橘子未被感染，说明它们被孤立了，返回-1
	if freshCount > 0 {
		return -1
	}
	return minutes
}
