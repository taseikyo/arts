/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-09-11 14:35:15
 * @link    github.com/taseikyo
 */

func totalNumbers(digits []int) int {
	set := map[int]bool{}
	for i, a := range digits { // 个位数
		if a%2 > 0 {
			continue
		}
		for j, b := range digits { // 十位数
			if j == i {
				continue
			}
			for k, c := range digits { // 百位数
				if c == 0 || k == i || k == j {
					continue
				}
				set[c*100+b*10+a] = true
			}
		}
	}
	return len(set)
}
