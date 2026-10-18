/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-23 01:04:45
 * @link    github.com/taseikyo
 */

func twoSum(nums []int, target int) []int {
	if nums[0] > target {
		return nil
	}

	// 初始化双指针，i 指向数组头部，j 指向数组尾部
	i, j := 0, len(nums)-1

	for i < j {
		sum := nums[i] + nums[j]
		if sum == target {
			// 找到目标，直接返回
			return []int{nums[i], nums[j]}
		} else if sum < target {
			// 和太小，左指针右移
			i++
		} else {
			// 和太大，右指针左移
			j--
		}
	}

	// 题目保证有解，此处仅为编译通过
	return []int{}
}
