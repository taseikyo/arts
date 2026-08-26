/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-22 00:18:46
 * @link    github.com/taseikyo
 */

func countNumbers(cnt int) []int {
	if cnt < 1 {
		return nil
	}
	size := 10
	for i := 1; i < cnt; i++ {
		size = size * 10
	}

	res := make([]int, 0, size-1)
	for i := 1; i < size; i++ {
		res = append(res, i)
	}

	return res
}
