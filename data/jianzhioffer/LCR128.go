/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-21 15:59:38
 * @link    github.com/taseikyo
 */

func inventoryManagement(stock []int) int {
	if len(stock) < 2 {
		return stock[0]
	}
	l, r := 0, len(stock)-1

	for l <= r {
		if l == r {
			return stock[r]
		}
		mid := (r-l)/2 + l
		if stock[mid] > stock[r] {
			l = mid + 1
		} else if stock[mid] < stock[r] {
			r = mid
		} else {
			r--
		}
	}

	return stock[r]
}
