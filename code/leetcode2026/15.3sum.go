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

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		if n1 > 0 {
			break
		}
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				for l < r && nums[l] == n2 {
					l++
				}
				for l < r && nums[r] == n3 {
					r--
				}
			} else if n1+n2+n3 < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}
