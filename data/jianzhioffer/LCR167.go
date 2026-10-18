/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-24 01:25:40
 * @link    github.com/taseikyo
 */

func dismantlingAction(s string) int {
	// charIndex 记录每个字符上一次出现的位置
	charIndex := make(map[byte]int)
	maxLen := 0
	// left 指针指向当前窗口的起始位置
	left := 0

	for i := 0; i < len(s); i++ {
		ch := s[i]

		// 如果当前字符之前出现过，并且它还在当前的窗口内
		if lastIdx, ok := charIndex[ch]; ok && lastIdx >= left {
			// 更新 left 指针，跳到重复字符的下一个位置，将其移出窗口
			left = lastIdx + 1
		}

		// 更新当前字符的最新位置
		charIndex[ch] = i

		// 计算当前窗口长度并更新最大长度
		currentLen := i - left + 1
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}
	return maxLen
}
