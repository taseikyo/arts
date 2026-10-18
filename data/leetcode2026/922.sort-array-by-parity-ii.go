/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-31 15:45:39
 * @link    github.com/taseikyo
 */

func sortArrayByParityII(nums []int) []int {
	i, j, n := 0, 1, len(nums)
	for i < n && j < n {
		for i < n && nums[i]%2 == 0 {
			i += 2
		}
		for j < n && nums[j]%2 != 0 {
			j += 2
		}
		if i < n && j < n {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	return nums
}
