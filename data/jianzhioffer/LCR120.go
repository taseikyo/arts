/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-22 00:13:55
 * @link    github.com/taseikyo
 */

func findRepeatDocument(documents []int) int {
	cache := make(map[int]bool, 0)
	for _, v := range documents {
		if ok, e := cache[v]; ok && e {
			return v
		}
		cache[v] = true
	}

	return -1
}
