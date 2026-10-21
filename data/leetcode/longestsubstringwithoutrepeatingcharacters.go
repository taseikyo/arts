/*
* @Date:   2025-03-02 19:00:44
* @Author: Lewis Tian (taseikyo@gmail.com)
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
