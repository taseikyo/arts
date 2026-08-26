/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-21 14:41:24
 * @link    github.com/taseikyo
 */

func rob(nums []int) int {
	n := len(nums)
	dp := make([][]int, n)
	for idx := range dp {
		dp[idx] = make([]int, 2)
	}

	// dp[i][0] 不偷，dp[i][1] 偷
	dp[0][1] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]
	}

	return max(dp[n-1][0], dp[n-1][1])
}

func rob(nums []int) int {
    // 初始化第0间房子的状态
    notRob := 0     // dp[0][0]
    rob := nums[0]  // dp[0][1]

    for i := 1; i < len(nums); i++ {
        // 计算当前房子的两个状态
        newNotRob := max(notRob, rob)          // 当前不偷 = 上一间偷/不偷的最大值
        newRob := notRob + nums[i]             // 当前偷 = 上一间不偷 + 当前金额
        // 更新状态
        notRob, rob = newNotRob, newRob
    }
    return max(notRob, rob)
}
