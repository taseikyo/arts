/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-21 00:13:03
 * @link    github.com/taseikyo
 */

func subsets(nums []int) [][]int {
	var res [][]int
	var path []int

	var backtrack func(start int)
	backtrack = func(start int) {
		// 关键：将当前路径的拷贝加入结果集（每个节点都是一个子集）
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)

		// 从 start 开始遍历，保证不会出现重复子集
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i]) // 做选择
			backtrack(i + 1)             // 递归
			path = path[:len(path)-1]    // 撤销选择（回溯）
		}
	}

	backtrack(0)
	return res
}
