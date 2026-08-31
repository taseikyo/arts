/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-27 01:30:10
 * @link    github.com/taseikyo
 */

func fib(n int) int {
	if n < 2 {
		return n
	}
	a, b := 0, 1
	const mod = 1000000007
	for i := 2; i <= n; i++ {
		a, b = b, (a+b)%mod
	}
	return b
}
