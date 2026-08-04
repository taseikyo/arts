/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-02 16:30:06
 * @link    github.com/taseikyo
 */

func lengthOfLongestSubstring(s string) (ans int) {
	for i, j := 0, 0; j < len(s); j++ {
		for strings.Contains(s[i:j], string(s[j])) {
			i++
		}
		ans = max(ans, j-i+1)
	}
	return
}
