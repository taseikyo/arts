/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-07-13 14:15:00
 * @link    github.com/taseikyo
 */

func lengthOfLastWord(s string) int {
	var idx, count int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			continue
		}
		idx = i
		break
	}
	for i := idx; i >= 0; i-- {
		if s[i] == ' ' {
			break
		}
		count++
	}

	return count
}
