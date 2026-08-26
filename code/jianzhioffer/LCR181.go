/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-24 01:08:15
 * @link    github.com/taseikyo
 */

func reverseMessage(s string) string {
	// 1. 去除首尾空格
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return ""
	}

	var res []string
	left, right := len(s)-1, len(s)-1

	for left >= 0 {
		// 2. 找到当前单词的起始位置 (left)
		for left >= 0 && s[left] != ' ' {
			left--
		}
		// 此时 left 指向空格或 -1，单词范围是 [left+1, right]
		res = append(res, s[left+1:right+1])

		// 3. 跳过单词间的空格，准备寻找下一个单词
		for left >= 0 && s[left] == ' ' {
			left--
		}
		// 更新右指针到下一个单词的末尾
		right = left
	}

	// 4. 用单个空格连接所有单词
	return strings.Join(res, " ")
}
