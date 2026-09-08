/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-09-04 16:21:31
 * @link    github.com/taseikyo
 */

import "math"

func firstStableIndex(nums []int, k int) int {
	n := len(nums)
	maxNums, minNums := make([]int, n), make([]int, n)
	maxNums[0] = nums[0]
	minNums[n-1] = nums[n-1]
	for i := 1; i < n; i++ {
		maxNums[i] = maxNums[i-1]
		if nums[i] > maxNums[i-1] {
			maxNums[i] = nums[i]
		}
	}
	for i := n - 2; i >= 0; i-- {
		minNums[i] = minNums[i+1]
		if nums[i] < minNums[i+1] {
			minNums[i] = nums[i]
		}
	}

	res := math.MaxInt
	for i := 0; i < n; i++ {
		if maxNums[i]-minNums[i] > k {
			continue
		}
		if res > maxNums[i]-minNums[i] {
			return i
		}
	}

	return -1
}
