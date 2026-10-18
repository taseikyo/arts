/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-21 00:35:59
 * @link    github.com/taseikyo
 */

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int // 存放最终结果
	var path []int  // 存放当前搜索路径

	// 1. 排序是为了方便剪枝，不是必须的
	sort.Ints(candidates)

	// 2. 定义 DFS 函数（闭包）
	var dfs func(start, sum int)
	dfs = func(start, sum int) {
		// 3. 终止条件：找到一个有效组合
		if sum == target {
			// 关键：拷贝一份 path，因为后续回溯会修改它
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		// 4. 剪枝：如果当前和已经超过目标，直接返回
		if sum > target {
			return
		}

		// 5. 遍历选择列表
		//    从 start 开始，保证每个元素可以被重复选择
		for i := start; i < len(candidates); i++ {
			// 做选择
			path = append(path, candidates[i])
			// 递归：注意这里的 start 参数是 i，而不是 i+1
			// 这表示下一层仍然可以从 candidates[i] 开始选择，实现无限次使用
			dfs(i, sum+candidates[i])
			// 撤销选择（回溯）
			path = path[:len(path)-1]
		}
	}

	dfs(0, 0)
	return res
}
