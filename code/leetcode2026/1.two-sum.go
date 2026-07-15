/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-07-12 16:18:06
 * @link    github.com/taseikyo
 */

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int)
	res := make([]int, 0)
	for i, num := range nums {
		val := target - num
		if j, ok := cache[val]; ok {
			res = append(res, []int{j, i}...)
			break
		}
		cache[num] = i
	}
	return res
}
