/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-02 21:00:06
 * @link    github.com/taseikyo
 */

func findAnagrams(s string, p string) []int {
	// 先判断特殊情况，如果 S 比 P 小，就直接返回空
	if len(s) < len(p) {
		return nil
	}
	need := [26]int{}
	for i := range p {
		need[p[i]-'a']++
	}
	window := [26]int{}
	left := 0
	res := []int{}
	for right := 0; right < len(s); right++ {
		window[s[right]-'a']++
		if right-left+1 == len(p) {
			if window == need {
				res = append(res, left)
			}
			window[s[left]-'a']--
			left++
		}
	}
	return res
}
