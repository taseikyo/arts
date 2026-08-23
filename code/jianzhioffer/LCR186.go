/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-23 21:16:36
 * @link    github.com/taseikyo
 */

import "sort"

func checkDynasty(nums []int) bool {
	sort.Ints(nums)

	// 统计大小王(0)的个数
	joker := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			joker++
		}
	}

	// 从第一个非0元素开始检查重复和最大差值
	// 第一个非0元素的索引就是 joker
	for i := joker; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] { // 有重复，不是顺子
			return false
		}
	}

	// 最大值 nums[4] - 最小值 nums[joker] < 5
	return nums[4]-nums[joker] < 5
}
