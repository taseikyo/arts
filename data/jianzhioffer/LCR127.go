/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-27 01:34:52
 * @link    github.com/taseikyo
 */

func trainWays(num int) int {
	if num == 0 {
		return 1
	}
	if num <= 2 {
		return num
	}

	a, b := 1, 2
	const mod = 1000000007
	for i := 3; i <= num; i++ {
		a, b = b, (a+b)%mod
	}
	return b
}
