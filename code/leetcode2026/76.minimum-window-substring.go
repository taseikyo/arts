/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-04 18:20:06
 * @link    github.com/taseikyo
 */

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	need, window := make(map[byte]int, 0), make(map[byte]int, 0)
	left, right, start, minLen := 0, 0, 0, 0
	match := 0

	for i, _ := range t {
		need[t[i]]++
	}
	for right < len(s) {
		ch := s[right]
		if v, ok := need[ch]; ok {
			window[ch]++
			if window[ch] == v {
				match++
			}
		}
		// 当前子串正好匹配t
		for match == len(need) {
			if minLen == 0 || minLen > right-left+1 {
				minLen = right - left + 1
				start = left
			}
			ch := s[left]
			if v, ok := need[ch]; ok {
				window[ch]--
				if window[ch] < v {
					match--
				}
			}
			left++
		}
		right++
	}

	return s[start : start+minLen]
}
