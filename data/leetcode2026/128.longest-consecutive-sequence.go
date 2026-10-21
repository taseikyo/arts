/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-07-23 00:03:19
 * @link    github.com/taseikyo
 */

func longestConsecutive(nums []int) int {
	cache := make(map[int]int)
	res := 0
	for _, num := range nums {
		if _, ok := cache[num]; !ok {
			prev := cache[num-1]
			next := cache[num+1]

			cache[num] += prev + next + 1
			cache[num-prev] = cache[num]
			cache[num+next] = cache[num]

			if cache[num] > res {
				res = cache[num]
			}
		}
	}

	return res
}
