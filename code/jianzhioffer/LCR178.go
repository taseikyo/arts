/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-27 01:24:28
 * @link    github.com/taseikyo
 */

func trainingPlan(nums []int) int {
	res := 0
	// 遍历 32 位（int 类型）
	for i := 0; i < 32; i++ {
		count := 0
		// 统计所有数字在第 i 位上 1 的个数
		for _, num := range nums {
			if (num>>i)&1 == 1 {
				count++
			}
		}
		// 如果 count 不是 3 的倍数，说明目标数字在这一位上是 1
		if count%3 != 0 {
			res |= 1 << i
		}
	}
	return res
}
