/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-23 20:35:03
 * @link    github.com/taseikyo
 */

func fileCombination(target int) [][]int {
	var res [][]int
	left, right := 1, 2
	sum := left + right

	// 因为序列至少包含两个数，所以 left 最大不会超过 target/2
	for left < right && left <= target/2 {
		if sum < target {
			// 窗口和太小，右指针右移扩大窗口
			right++
			sum += right
		} else if sum > target {
			// 窗口和太大，左指针右移缩小窗口
			sum -= left
			left++
		} else {
			// 找到符合条件的序列
			var seq []int
			for i := left; i <= right; i++ {
				seq = append(seq, i)
			}
			res = append(res, seq)
			// 继续寻找，将左指针右移
			sum -= left
			left++
		}
	}
	return res
}
