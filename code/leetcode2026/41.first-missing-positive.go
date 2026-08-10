/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-10 18:05:03
 * @link    github.com/taseikyo
 */

func firstMissingPositive(nums []int) int {
	minV := math.MaxInt
	maxV := 0
	cache := make(map[int]int, 0)
	for _, v := range nums {
		if v <= 0 {
			continue
		}
		cache[v] = 1
		minV = min(minV, v)
		maxV = max(maxV, v)
	}
	if minV > 1 {
		return 1
	}
	for i := 1; i < maxV; i++ {
		if _, ok := cache[i]; !ok {
			return i
		}
	}

	return maxV + 1
}
