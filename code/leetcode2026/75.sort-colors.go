/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-20 01:16:24
 * @link    github.com/taseikyo
 */

func sortColors(nums []int) {
	start, end, cur := 0, len(nums)-1, 0
	for cur <= end {
		if nums[cur] == 1 {
			cur++
		} else if nums[cur] == 2 {
			nums[cur], nums[end] = nums[end], nums[cur]
			end--
		} else if nums[cur] == 0 {
			nums[cur], nums[start] = nums[start], nums[cur]
			start++
            cur++
		}
	}
}
