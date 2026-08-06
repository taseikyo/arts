/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-06 22:30:06
 * @link    github.com/taseikyo
 */


func maxSubArray(nums []int) int {
    res := nums[0]
    for i := 0; i < len(nums); i++ {
        tmp := 0
        for j := i; j < len(nums); j++ {
            tmp += nums[j]
            res = max(res, tmp)
        }
    }
    return res
}

func maxSubArray(nums []int) int {
	ans := math.MinInt
	minPreSum := 0
	preSum := 0
	for _, x := range nums {
		preSum += x                        // 当前的前缀和
		ans = max(ans, preSum-minPreSum)   // 减去前缀和的最小值
		minPreSum = min(minPreSum, preSum) // 维护前缀和的最小值
	}
	return ans
}

func maxSubArray(nums []int) int {
    f := make([]int, len(nums))
    f[0] = nums[0]
    for i := 1; i < len(nums); i++ {
        f[i] = max(f[i-1], 0) + nums[i]
    }
    return slices.Max(f)
}
