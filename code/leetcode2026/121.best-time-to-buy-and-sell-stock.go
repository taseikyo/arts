/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-20 16:53:55
 * @link    github.com/taseikyo
 */

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	// minPrice: 遍历过程中遇到的历史最低价格
	// maxProfit: 当前能获取的最大利润
	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		// 如果当前价格低于历史最低，更新最低价
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			// 否则计算以当前价格卖出的利润，并更新最大利润
			profit := prices[i] - minPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}
