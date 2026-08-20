/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-20 17:25:24
 * @link    github.com/taseikyo
 */

func jump(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	jumps := 0      // 跳跃次数
	currentEnd := 0 // 当前跳跃能到达的边界
	farthest := 0   // 在 currentEnd 范围内，能到达的最远位置

	// 关键：不需要遍历最后一个元素
	for i := 0; i < n-1; i++ {
		// 1. 更新从当前位置能跳到的最远位置
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}

		// 2. 如果到达了当前跳跃的边界
		if i == currentEnd {
			jumps++               // 必须再跳一次
			currentEnd = farthest // 更新新的跳跃边界

			// 如果新的边界已经能覆盖终点，可以提前结束
			if currentEnd >= n-1 {
				break
			}
		}
	}
	return jumps
}

func jump(nums []int) int {
    n := len(nums)
    dp := make([]int, n)
    // dp[n-1] 默认为 0，表示终点到终点不需要步数

    for i := n - 2; i >= 0; i-- {
        minStep := int(1e9)
        // 检查从 i 能跳到的所有位置 j
        for j := i + 1; j <= i+nums[i] && j < n; j++ {
            if dp[j] < minStep {
                minStep = dp[j]
            }
        }
        dp[i] = minStep + 1
    }
    return dp[0]
}
