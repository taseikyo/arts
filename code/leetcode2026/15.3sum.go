/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-07-25 22:06:52
 * @link    github.com/taseikyo
 */

func threeSum(nums []int) [][]int {
	cache := make(map[int]int)
	for idx, v := range nums {
		cache[v] = idx
	}

	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			target := 0 - nums[i] - nums[j]
			if idx, ok := cache[target]; ok && idx != i && idx != j {
				tmp := []int{nums[i], nums[j], target}
				sort.Ints(tmp)
				flag := false
				for _, old := range res {
					if old[0] == tmp[0] && old[1] == tmp[1] && old[2] == tmp[2] {
						flag = true
						break
					}
				}
				if !flag {
					res = append(res, tmp)
				}
			}
		}
	}

	return res
}
