/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-07-25 21:43:09
 * @link    github.com/taseikyo
 */

func moveZeroes(nums []int) {
	a := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if nums[a] == 0 {
				nums[i], nums[a] = nums[a], nums[i]
			}
			a += 1
		}
	}
}
