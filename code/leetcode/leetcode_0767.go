/**
 * @date    2020-11-30 11:48:56
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @link 	github.com/taseikyo
 */

type pair struct {
	char  rune
	count int
}

func reorganizeString(S string) string {
	// a-z
	bucket := make([]pair, 26)
	maxCount := 0
	// 统计各个字符出现次数
	for _, char := range S {
		offset := char - 'a'
		if bucket[offset].count == 0 {
			bucket[offset].char = char
		}
		bucket[offset].count++
		if bucket[offset].count > maxCount {
			maxCount = bucket[offset].count
		}
	}
	if maxCount > (len(S)+1)/2 {
		// 超过一半，无解
		return ""
	}
	// 出现次数从大到小排序
	sort.Slice(bucket, func(i, j int) bool {
		return bucket[i].count >= bucket[j].count
	})

	// 从最多的字符开始按照下标先偶后奇填
	chars, index := make([]rune, len(S)), 0
	for i := 0; i < 2; i++ {
		for j := i; j < len(S); j += 2 {
			if bucket[index].count == 0 {
				// 消耗完当前字符
				index++
			}
			chars[j] = bucket[index].char
			bucket[index].count--
		}
	}
	return string(chars)
}