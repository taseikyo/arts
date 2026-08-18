/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-18 17:36:51
 * @link    github.com/taseikyo
 */

func findMin(nums []int) int {
    l, r := 0, len(nums)-1
    for l < r {
        mid := (r-l)/2 + l
        if nums[mid] > nums[r] {
            l = mid + 1
        } else {
            r = mid
        }
    }

    return nums[r]
}
