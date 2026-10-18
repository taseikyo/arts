/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-07-25 22:30:12
 * @link    github.com/taseikyo
 */

func maxArea(height []int) int {
	count, left, right := 0, 0, len(height)-1
	for left < right {
		count = max(count, (right-left)*min(height[left], height[right]))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return count
}
