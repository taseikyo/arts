/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-23 00:41:59
 * @link    github.com/taseikyo
 */

import (
	"sort"
	"strconv"
	"strings"
)

func crackPassword(nums []int) string {
	// 1. 将整数切片转换为字符串切片
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}

	// 2. 使用自定义规则排序
	sort.Slice(strs, func(i, j int) bool {
		// 核心：比较两种拼接方式的结果
		return strs[i]+strs[j] < strs[j]+strs[i]
	})

	// 3. 将排序后的字符串拼接起来
	var sb strings.Builder
	for _, s := range strs {
		sb.WriteString(s)
	}

	return sb.String()
}
