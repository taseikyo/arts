/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-20 17:07:28
 * @link    github.com/taseikyo
 */

func canJump(nums []int) bool {
	maxReach := 0
	for i := 0; i < len(nums); i++ {
		// 如果当前位置已经超出了能到达的最远距离，说明无法到达
		if i > maxReach {
			return false
		}
		// 更新能到达的最远距离
		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
		}
	}
	return true
}

func canJump(nums []int) bool {
    lastPos := len(nums) - 1
    for i := len(nums) - 2; i >= 0; i-- {
        if i+nums[i] >= lastPos {
            lastPos = i
        }
    }
    return lastPos == 0
}

func canJump(nums []int) bool {
    dp := make([]int, len(nums))
    for i := 1; i < len(nums); i++ {
        // 当前位置的剩余步数 = max(上一步剩余步数, 上一步能跳的步数) - 1
        dp[i] = max(dp[i-1], nums[i-1]) - 1
        if dp[i] < 0 {
            return false
        }
    }
    return true
}
