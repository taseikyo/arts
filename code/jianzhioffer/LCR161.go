/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-27 01:40:05
 * @link    github.com/taseikyo
 */

func maxSales(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	// dp[i] 表示以 nums[i] 结尾的连续子数组的最大和
	dp := make([]int, n)
	dp[0] = nums[0]
	// 最终答案初始化为第一个元素
	ans := dp[0]

	for i := 1; i < n; i++ {
		// 状态转移方程：要么自立门户，要么加入之前的子数组
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		// 更新全局最大值
		ans = max(ans, dp[i])
	}
	return ans
}
