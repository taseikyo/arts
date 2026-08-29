/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-29 15:19:43
 * @link    github.com/taseikyo
 */

func removeDuplicates(nums []int) int {
	i, j := 0, 0
	for j < len(nums) {
		for j < len(nums) && nums[j] == nums[i] {
			j++
		}
		if j == len(nums) {
			break
		}
		i++
		nums[i] = nums[j]
	}

	return i + 1
}
