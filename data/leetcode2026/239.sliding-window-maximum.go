/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-04 17:18:06
 * @link    github.com/taseikyo
 */

// maxSlidingWindow 函数用于求解滑动窗口中的最大值
// 参数 nums 是一个整数切片，代表输入的数组
// 参数 k 是一个整数，代表滑动窗口的大小
func maxSlidingWindow(nums []int, k int) []int {
	// 定义一个栈（这里用切片模拟），用于存储当前窗口内可能是最大值的元素的索引
	var stack []int
	// 定义一个结果切片，用于存储每个滑动窗口的最大值
	var res []int
	// 遍历输入数组 nums
	for i, v := range nums {
		// 当栈不为空，并且当前元素 v 大于等于栈顶元素对应的数组值时
		for len(stack) > 0 && v >= nums[stack[len(stack)-1]] {
			// 弹出栈顶元素，因为栈顶元素对应的数组值不可能是当前或后续窗口的最大值了
			stack = stack[:len(stack)-1]
		}
		// 将当前元素的索引 i 压入栈中
		stack = append(stack, i)
		// 如果当前索引 i 减去窗口大小 k 再加 1 大于栈底元素（即栈中第一个元素）的索引
		if i-k+1 > stack[0] {
			// 说明栈底元素已经不在当前窗口内了，将其从栈中移除
			stack = stack[1:]
		}
		// 如果当前索引 i 加 1 大于等于窗口大小 k，说明已经形成了一个完整的窗口
		if i+1 >= k {
			// 栈底元素对应的数组值就是当前窗口的最大值，将其添加到结果切片中
			res = append(res, nums[stack[0]])
		}
	}
	// 返回结果切片
	return res
}
