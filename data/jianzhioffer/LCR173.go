/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-23 01:05:45
 * @link    github.com/taseikyo
 */

func takeAttendance(records []int) int {
	l, r := 0, len(records)-1
	for l <= r {
		mid := (r-l)/2 + l
		if records[mid] > mid {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

    return l
}
