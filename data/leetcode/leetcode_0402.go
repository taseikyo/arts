/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2020-11-17 20:44:05
 * @link    github.com/taseikyo
 */

func removeKdigits(num string, k int) string {
	stack := []byte{}
	for i := range num {
		digit := num[i]
		for k > 0 && len(stack) > 0 && digit < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, digit)
	}
	stack = stack[:len(stack)-k]
	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		ans = "0"
	}
	return ans
}
