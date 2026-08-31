/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-27 16:20:57
 * @link    github.com/taseikyo
 */

func findNthDigit(n int) int {
	// 1. 确定位数 k
	// 使用 int64 防止乘法溢出
	var digitLen int64 = 1 // 当前位数，从1开始
	var count int64 = 9    // 当前位数区间有多少个数字
	var start int64 = 1    // 当前位数区间的起始数字

	n64 := int64(n)

	// 不断减去当前区间的总位数，直到 n 落在某个区间
	for n64 > count*digitLen {
		n64 -= count * digitLen
		digitLen++
		count *= 10
		start *= 10
	}

	// 2. 确定具体数字
	// 在区间内偏移量（0-based）
	offset := (n64 - 1) / digitLen
	num := start + offset

	// 3. 确定数字中的位置
	pos := (n64 - 1) % digitLen // 0 表示最高位

	// 将数字转为字符串，取第 pos 位
	str := strconv.FormatInt(num, 10)
	return int(str[pos] - '0')
}
