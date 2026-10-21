/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-02 16:30:06
 * @link    github.com/taseikyo
 */

func trap(height []int) int {
	length := len(height)
	leftMax := make([]int, length)
	rightMax := make([]int, length)
	if length < 2 {
		return 0
	}

	leftMax[0] = height[0]
	rightMax[length-1] = height[length-1]
	for i := 1; i < length; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	for i := length - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	total := 0
	for i := 0; i < length; i++ {
		total += min(leftMax[i], rightMax[i]) - height[i]
	}

	return total
}

func trap(height []int) int {
	if len(height) < 2 {
		return 0
	}
	leftMax, rightMax, total := 0, 0, 0
	left, right := 0, len(height)-1
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				total += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				total += rightMax - height[right]
			}
			right--
		}
	}

	return total
}
