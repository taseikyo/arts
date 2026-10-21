/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-20 16:26:12
 * @link    github.com/taseikyo
 */

func findDuplicate(nums []int) int {
    // 1. 寻找相遇点
    // 初始化：慢指针走一步，快指针走两步
    slow := nums[0]
    fast := nums[nums[0]]

    for slow != fast {
        slow = nums[slow]           // 慢指针每次走一步
        fast = nums[nums[fast]]     // 快指针每次走两步
    }

    // 2. 寻找环的入口
    fast = 0 // 将快指针重置到起点
    for slow != fast {
        slow = nums[slow] // 慢指针每次走一步
        fast = nums[fast] // 快指针每次走一步
    }

    // 再次相遇的点就是重复的数字
    return slow
}
