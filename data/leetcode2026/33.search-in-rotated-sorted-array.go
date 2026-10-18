/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-18 17:26:45
 * @link    github.com/taseikyo
 */

func search(nums []int, target int) int {
    left, right := 0, len(nums)-1

    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            return mid
        }

        // 判断左半部分是否有序
        if nums[left] <= nums[mid] {
            // 左半部分有序，检查 target 是否在左半部分
            if nums[left] <= target && target < nums[mid] {
                right = mid - 1 // 在左半部分继续搜索
            } else {
                left = mid + 1 // 在右半部分继续搜索
            }
        } else {
            // 右半部分有序，检查 target 是否在右半部分
            if nums[mid] < target && target <= nums[right] {
                left = mid + 1 // 在右半部分继续搜索
            } else {
                right = mid - 1 // 在左半部分继续搜索
            }
        }
    }

    return -1
}
