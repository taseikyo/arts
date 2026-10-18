/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-23 00:20:57
 * @link    github.com/taseikyo
 */

func nthUglyNumber(n int) int {
	// dp[i] 表示第 i+1 个丑数
	dp := make([]int, n)
	dp[0] = 1 // 第一个丑数是 1[reference:12]

	// 初始化三个指针，都指向第一个丑数
	p2, p3, p5 := 0, 0, 0

	for i := 1; i < n; i++ {
		// 计算下一个丑数的三个候选值[reference:13]
		next2 := dp[p2] * 2
		next3 := dp[p3] * 3
		next5 := dp[p5] * 5

		// 取最小值作为下一个丑数[reference:14]
		dp[i] = min(next2, min(next3, next5))

		// 判断是由哪个指针贡献的，并将对应指针后移[reference:15]
		// 使用 if 而非 if-else，是为了处理 next2 == next3 这种重复情况，可以同时后移，达到去重效果[reference:16]
		if dp[i] == next2 {
			p2++
		}
		if dp[i] == next3 {
			p3++
		}
		if dp[i] == next5 {
			p5++
		}
	}

	return dp[n-1]
}
