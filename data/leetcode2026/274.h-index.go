/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-30 16:24:05
 * @link    github.com/taseikyo
 */

import "sort"

func hIndex(citations []int) int {
	// 1. 升序排序
	sort.Ints(citations)
	n := len(citations)

	// 2. 从后往前找 h
	// 当 citations[n-h] >= h 时，说明有至少 h 篇论文的引用次数 >= h[reference:10]
	for h := n; h > 0; h-- {
		if citations[n-h] >= h {
			return h
		}
	}
	return 0
}
