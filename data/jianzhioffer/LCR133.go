/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-26 16:36:45
 * @link    github.com/taseikyo
 */

func hammingWeight(num uint32) int {
	res := 0
	for num > 0 {
		if num&1 == 1 {
			res++
		}
		num = num >> 1
	}
	return res
}
