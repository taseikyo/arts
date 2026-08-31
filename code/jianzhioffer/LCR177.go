/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-27 01:16:44
 * @link    github.com/taseikyo
 */

func sockCollocation(nums []int) []int {
	// 1. 全员异或，得到 a ^ b
	xor := 0
	for _, num := range nums {
		xor ^= num
	}

	// 2. 找到 xor 中最低位的 1 作为分组依据
	// 使用 xor & -xor 可以快速得到最低位 1
	mask := xor & -xor

	// 3. 分组异或
	num1, num2 := 0, 0
	for _, num := range nums {
		if num&mask == 0 {
			num1 ^= num
		} else {
			num2 ^= num
		}
	}

	return []int{num1, num2}
}
