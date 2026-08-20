/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-20 01:06:29
 * @link    github.com/taseikyo
 */

func singleNumber(nums []int) int {
	x := nums[0]
	for i := 1; i < len(nums); i++ {
		x ^= nums[i]
	}

	return x
}
