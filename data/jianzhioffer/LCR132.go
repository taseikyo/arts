/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-22 00:24:45
 * @link    github.com/taseikyo
 */

func cuttingBamboo(n int) int {
	const mod = 1000000007

	// 1. 处理基础情况
	if n <= 3 {
		return n - 1
	}

	// 2. 贪心：尽可能多地剪出长度为 3 的段
	res := 1
	for n > 4 {
		res = (res * 3) % mod
		n -= 3
	}

	// 3. 最后剩下的长度是 2, 3, 或 4
	//    当 n=4 时，对应余数为 1 的情况（4 = 3 + 1，合并为 2+2）
	//    当 n=3 时，对应余数为 0 的情况
	//    当 n=2 时，对应余数为 2 的情况
	return (res * n) % mod
}
