/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-24 00:58:25
 * @link    github.com/taseikyo
 */

func validNumber(s string) bool {
	// 1. 去除首尾空格
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return false
	}

	hasNum, hasDot, hasE := false, false, false

	for i, ch := range s {
		if ch >= '0' && ch <= '9' {
			// 遇到数字，标记 hasNum
			hasNum = true
		} else if ch == 'e' || ch == 'E' {
			// e/E: 之前必须有数字，且不能重复出现
			if hasE || !hasNum {
				return false
			}
			hasE = true
			hasNum = false // 重置，因为e后面必须还有数字
		} else if ch == '+' || ch == '-' {
			// 正负号: 只能出现在开头，或紧跟在 e/E 后面
			if i != 0 && s[i-1] != 'e' && s[i-1] != 'E' {
				return false
			}
			// 注意：这里不需要设置标志，因为它的合法性由位置决定
		} else if ch == '.' {
			// 小数点: 只能出现一次，且不能在 e/E 之后出现
			if hasDot || hasE {
				return false
			}
			hasDot = true
		} else {
			// 其他任何字符都是非法的
			return false
		}
	}

	// 最终必须出现过至少一个数字
	return hasNum
}
