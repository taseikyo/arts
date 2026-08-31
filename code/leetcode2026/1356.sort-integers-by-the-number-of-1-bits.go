/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-31 15:19:30
 * @link    github.com/taseikyo
 */

func sortByBits(arr []int) []int {
	bitsMap := make(map[int]int)
	for i, v := range arr {
		cnt := 0
		for v > 0 {
			if v&1 == 1 {
				cnt++
			}
			v = v >> 1
		}
		bitsMap[arr[i]] = cnt
	}
	slices.SortFunc(arr, func(a, b int) int {
		if bitsMap[a] == bitsMap[b] {
			return a - b
		}
		return bitsMap[a] - bitsMap[b]
	})

    return arr
}
