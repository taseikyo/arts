/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-03 20:58:06
 * @link    github.com/taseikyo
 */

func subarraySum(nums []int, k int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		tmp := 0
		for j := i; j < len(nums); j++ {
			tmp += nums[j]
			if tmp == k {
				res++
			}
		}
	}

	return res
}

func subarraySum(nums []int, k int) int {
	count := 0
	hash := map[int]int{0: 1}
	preSum := 0

	for i := 0; i < len(nums); i++ {
		preSum += nums[i]
		if hash[preSum-k] > 0 {
			count += hash[preSum-k]
		}
		hash[preSum]++
	}
	return count
}
