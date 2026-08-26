/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-24 14:54:19
 * @link    github.com/taseikyo
 */

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	// 1. 去除首尾空格
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	// 2. 处理符号（只允许出现在开头，且最多一次）
	minus := false
	i := 0
	if s[i] == '+' || s[i] == '-' {
		if s[i] == '-' {
			minus = true
		}
		i++
		// 如果只有一个符号，没有数字，返回 0
		if i == len(s) {
			return 0
		}
	}

	res := 0
	// 3. 遍历数字部分
	for ; i < len(s); i++ {
		ch := s[i]
		if ch < '0' || ch > '9' {
			break // 遇到非数字字符，停止转换
		}
		digit := int(ch - '0')

		// 4. 溢出截断（关键修改）
		// 如果 res 已经大于 (maxInt32 - digit) / 10，则下一步一定溢出
		if res > (math.MaxInt32-digit)/10 {
			if minus {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		}

		res = res*10 + digit
	}

	// 5. 应用符号
	if minus {
		res = -res
	}
	return res
}
