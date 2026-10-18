/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-18 17:01:32
 * @link    github.com/taseikyo
 */

func searchRange(nums []int, target int) []int {
	// 初始化结果为 -1
	first, last := -1, -1

	// 查找左边界
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			first = mid
			r = mid - 1 // 关键：继续向左搜索，看是否还有更靠前的 target
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	// 如果没找到左边界，说明 target 不存在，直接返回
	if first == -1 {
		return []int{-1, -1}
	}

	// 查找右边界
	l, r = 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			last = mid
			l = mid + 1 // 关键：继续向右搜索，看是否还有更靠后的 target
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return []int{first, last}
}
