/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-29 15:33:43
 * @link    github.com/taseikyo
 */

func removeDuplicates(nums []int) int {
	n := len(nums)
	// 如果数组长度小于等于2，则一定符合要求，直接返回原长度[reference:15]
	if n <= 2 {
		return n
	}

	// 初始化慢指针为2，因为前两个元素无论如何都是符合要求的[reference:16][reference:17]
	slow := 2
	// 快指针从索引2开始遍历[reference:18]
	for fast := 2; fast < n; fast++ {
		// 检查当前元素是否与慢指针前两个位置的元素相同
		// 如果不相同，则当前元素需要被保留
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++ // 慢指针后移，指向下一个待放置位置
		}
		// 如果相同，说明已经有两个了，当前元素是多余的，直接跳过
	}
	// 循环结束，slow的值就是新数组的长度
	return slow
}
