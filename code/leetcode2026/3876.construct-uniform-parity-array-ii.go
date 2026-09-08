/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-09-04 00:03:59
 * @link    github.com/taseikyo
 */

func uniformArray(nums1 []int) bool {
	// 只有一种情况是false -> 最小值是偶数 && 数组中出现了奇数
	// 其余情况都是true
	if slices.Min(nums1)%2 == 0 {
		for _, v := range nums1 {
			if v%2 != 0 {
				return false
			}
		}
	}

	return true
}
