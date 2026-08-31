/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-27 01:24:51
 * @link    github.com/taseikyo
 */

func encryptionCalculate(a int, b int) int {
	// 当进位 b 不为 0 时，持续循环
	for b != 0 {
		// 1. 计算进位
		carry := (a & b) << 1
		// 2. 计算无进位和
		a = a ^ b
		// 3. 将进位赋值给 b，进入下一轮
		b = carry
	}
	// 当 b == 0 时，a 即为最终结果
	return a
}
