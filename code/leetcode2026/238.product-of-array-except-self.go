/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-10 17:23:03
 * @link    github.com/taseikyo
 */

func productExceptSelf(nums []int) []int {
	length := len(nums)
	left, right, res := make([]int, length), make([]int, length), make([]int, length)
	left[0] = 1
	right[length-1] = 1
	for i := 1; i < length; i++ {
		left[i] = nums[i-1] * left[i-1]
	}
	for i := length - 2; i >= 0; i-- {
		right[i] = nums[i+1] * right[i+1]
	}

	for i := 0; i < length; i++ {
		res[i] = left[i] * right[i]
	}

	return res
}

func productExceptSelf(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	res[0] = 1
	for i := 1; i < length; i++ {
		res[i] = nums[i-1] * res[i-1]
	}
	tmp := nums[length-1]
	for i := length - 2; i >= 0; i-- {
		res[i] = tmp * res[i]
		tmp = tmp * nums[i]
	}

	return res
}
