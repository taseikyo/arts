/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-21 00:07:52
 * @link    github.com/taseikyo
 */

func permute(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)

	var backtrack func(int)
	backtrack = func(idx int) {
		// 终止条件：所有位置都已确定，得到一个完整排列
		if idx == n {
			// 注意：需要拷贝一份 nums，因为后续交换会修改它
			tmp := make([]int, n)
			copy(tmp, nums)
			res = append(res, tmp)
			return
		}

		for i := idx; i < n; i++ {
			// 做选择：将 nums[i] 交换到 idx 位置
			nums[idx], nums[i] = nums[i], nums[idx]
			// 递归确定下一个位置
			backtrack(idx + 1)
			// 撤销选择（回溯）：交换回来，恢复原状
			nums[idx], nums[i] = nums[i], nums[idx]
		}
	}

	backtrack(0)
	return res
}
