/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2021-01-08 19:20:21
 * @link    github.com/taseikyo
 */

func reverse(nums []int, start, end int) {
	for start < end {
		temp := nums[start]
		nums[start] = nums[end]
		start++
		nums[end] = temp
		end--
	}
}

func rotate(nums []int, k int)  {
	n := len(nums)
	k %= n
	reverse(nums, 0, n - 1)
	reverse(nums, 0, k - 1)
	reverse(nums, k, n - 1)
}

// 最快的题解
func rotate1(nums []int, k int)  {
    k %= len(nums)
    ans := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
    nums = append(nums[:0], ans...)

}

