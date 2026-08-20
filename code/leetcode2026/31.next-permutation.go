/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-20 16:32:25
 * @link    github.com/taseikyo
 */

func nextPermutation(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	// 1. 从右向左，找第一个升序对 (nums[i] < nums[i+1])
	//    这个 nums[i] 就是待交换的“较小数”
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	// 2. 如果找到了升序对 (i >= 0)
	if i >= 0 {
		// 从右向左，找第一个大于 nums[i] 的数 nums[j]
		// 这个 nums[j] 就是待交换的“较大数”
		j := n - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		// 交换“较小数”和“较大数”
		nums[i], nums[j] = nums[j], nums[i]
	}

	// 3. 将 i 位置之后的所有元素反转
	//    - 如果 i >= 0，反转后使右侧变为最小排列
	//    - 如果 i == -1，反转整个数组，得到最小排列
	reverse(nums[i+1:])
}

// 反转切片
func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
