/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-21 15:21:39
 * @link    github.com/taseikyo
 */

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLen := 0, 1

	// 辅助函数：从中心向两边扩展，返回回文串的长度[reference:13]
	expandAroundCenter := func(left, right int) int {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		// 当循环退出时，有效回文串的范围是 [left+1, right-1]
		// 长度为 right - left - 1[reference:14]
		return right - left - 1
	}

	for i := 0; i < len(s); i++ {
		// 1. 奇数长度回文，中心为 s[i][reference:15]
		len1 := expandAroundCenter(i, i)
		// 2. 偶数长度回文，中心为 s[i] 和 s[i+1][reference:16]
		len2 := expandAroundCenter(i, i+1)

		curMaxLen := max(len1, len2)
		if curMaxLen > maxLen {
			maxLen = curMaxLen
			// 计算当前最长回文子串的起始位置[reference:17]
			start = i - (maxLen-1)/2
		}
	}

	return s[start : start+maxLen]
}
