/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-24 00:53:33
 * @link    github.com/taseikyo
 */

func goodsOrder(s string) []string {
	// 1. 将字符串转为字节数组并排序，使相同字符相邻
	bytes := []byte(s)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})

	res := []string{}
	path := []byte{}
	visited := make([]bool, len(bytes))

	// 2. 定义回溯函数
	var dfs func()
	dfs = func() {
		if len(path) == len(bytes) {
			res = append(res, string(path))
			return
		}

		for i := 0; i < len(bytes); i++ {
			// 剪枝条件：如果当前字符已被访问，跳过
			if visited[i] {
				continue
			}
			// 关键剪枝：如果当前字符和前一个字符相同，且前一个字符未被访问，则跳过
			// 这保证了相同字符在同一个位置只会被固定一次
			if i > 0 && bytes[i] == bytes[i-1] && !visited[i-1] {
				continue
			}

			// 做选择
			visited[i] = true
			path = append(path, bytes[i])

			dfs() // 递归

			// 撤销选择（回溯）
			path = path[:len(path)-1]
			visited[i] = false
		}
	}

	dfs()
	return res
}
