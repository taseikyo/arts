/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-24 01:10:56
 * @link    github.com/taseikyo
 */

func dismantlingAction(s string) byte {
	// 1. 创建哈希表
	cnt := make(map[byte]int)

	// 2. 第一次遍历：统计频率
	for i := 0; i < len(s); i++ {
		cnt[s[i]]++
	}

	// 3. 第二次遍历：按原字符串顺序查找第一个频率为1的字符
	for i := 0; i < len(s); i++ {
		if cnt[s[i]] == 1 {
			return s[i]
		}
	}

	// 4. 没有找到，返回单空格
	return ' '
}
