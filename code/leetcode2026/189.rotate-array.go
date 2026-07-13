/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-07-13 13:43:03
 * @link    github.com/taseikyo
 */


// Sol 1
func rotate(nums []int, k int) {
	if k == 0 {
		return
	}
    if k>len(nums) {
        k = k % len(nums)
    }

	res := make([]int, 0, k)
	for i := len(nums) - k; i < len(nums); i++ {
		res = append(res, nums[i])
	}
	for i := len(nums) - 1; i >= k; i-- {
		nums[i] = nums[i-k]
	}
	for i := 0; i < k; i++ {
		nums[i] = res[i]
	}
}


// Sol 2
func rotate(nums []int, k int) {
	if k == 0 {
		return
	}
	if k > len(nums) {
		k = k % len(nums)
	}
	for i := 0; i < k; i++ {
		rotateOne(nums)
	}
}

func rotateOne(nums []int) {
	res := nums[len(nums)-1]
	for i := len(nums) - 1; i >= 1; i-- {
		nums[i] = nums[i-1]
	}
	nums[0] = res
}

// Sol 3

func rotate(nums []int, k int) {
	if k == 0 {
		return
	}
	if k > len(nums) {
		k = k % len(nums)
	}
	reverse(nums, 0, len(nums)-1)
    reverse(nums, 0, k-1)
    reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, l, r int) {
	for l < r {
		tmp := nums[l]
		nums[l] = nums[r]
		nums[r] = tmp
		l++
		r--
	}
}
