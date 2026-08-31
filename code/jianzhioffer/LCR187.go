/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-27 15:03:59
 * @link    github.com/taseikyo
 */

func iceBreakingGame(n int, m int) int {
	// f(1) = 0
	f := 0
	// 从 i=2 开始递推到 i=n
	for i := 2; i <= n; i++ {
		// 套用递推公式 f(i) = (f(i-1) + m) % i
		f = (f + m) % i
	}
	return f
}
