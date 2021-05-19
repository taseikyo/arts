/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2020-11-22 21:47:12
 * @link    github.com/taseikyo
 */

func canCompleteCircuit(gas []int, cost []int) int {
	left := 0
	start := 0
	totalGas, totalCost := 0, 0
	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
		left += gas[i] - cost[i]
		if left < 0 {
			start = i + 1
			left = 0
		}
	}
	if totalGas < totalCost {
		return -1
	}
	// 总加油>=总耗油，必然有解。当遍历结束时，最新的start指向成功的起点
	return start
}
