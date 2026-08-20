/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-20 17:35:22
 * @link    github.com/taseikyo
 */

func partitionLabels(s string) []int {
	// 1. 记录每个字符最后出现的位置
	// 因为题目只包含小写字母，使用长度为26的数组效率更高[reference:12][reference:13]
	lastPos := [26]int{}
	for i, ch := range s {
		lastPos[ch-'a'] = i
	}

	// 2. 遍历字符串，划分片段
	var res []int
	start, end := 0, 0
	for i, ch := range s {
		// 更新当前片段的结束边界
		if lastPos[ch-'a'] > end {
			end = lastPos[ch-'a']
		}

		// 如果当前位置就是片段的结束边界
		if i == end {
			// 记录片段长度
			res = append(res, end-start+1)
			// 开始寻找下一个片段
			start = end + 1
		}
	}
	return res
}
