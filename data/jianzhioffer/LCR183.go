/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-23 21:09:37
 * @link    github.com/taseikyo
 */

func maxAltitude(heights []int, limit int) []int {
	stack := []int{}
	res := []int{}

	for i, v := range heights {
		for len(stack) > 0 && v > heights[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
		if stack[0]+limit <= i {
			stack = stack[1:]
		}
		if i+1 >= limit {
			res = append(res, heights[stack[0]])
		}
	}

	return res
}
