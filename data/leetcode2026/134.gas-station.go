/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-31 16:04:48
 * @link    github.com/taseikyo
 */

func canCompleteCircuit(gas []int, cost []int) int {
	total := 0 // 总油量-总消耗，用于判断全局是否有解
	cur := 0   // 当前油箱的油量
	start := 0 // 记录可能的起点

	for i := 0; i < len(gas); i++ {
		diff := gas[i] - cost[i]
		total += diff
		cur += diff

		// 如果当前油量变为负数，说明从 start 到 i 都无法作为起点
		// 将起点设为下一个加油站 i+1，并重置当前油量
		if cur < 0 {
			start = i + 1
			cur = 0
		}
	}

	// 如果总油量小于总消耗，则无论从哪出发都无法完成环绕
	if total < 0 {
		return -1
	}
	return start
}
