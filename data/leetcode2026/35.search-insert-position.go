/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-18 16:38:29
 * @link    github.com/taseikyo
 */

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (r-l)/2 + l
        fmt.Println(m)
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	if l > r {
		return l
	}
	return r
}
