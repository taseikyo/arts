/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2020-11-15 19:41:58
 * @link    github.com/taseikyo
 */

func sortArrayByParityII(A []int) []int {
	ret := make([]int, len(A))
	even_idx, odd_idx := 0, 1
	for _, x := range(A)  {
		if x % 2 == 0 {
			ret[even_idx] = x
			even_idx += 2
		} else {
			ret[odd_idx] = x
			odd_idx += 2
		}
	}
	return ret
}
