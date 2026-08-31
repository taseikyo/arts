/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-30 16:17:31
 * @link    github.com/taseikyo
 */

func maxProfit(prices []int) int {
    profit := 0
    for i := 1; i < len(prices); i++ {
        // 如果今天比昨天价格高，就累加差价
        if prices[i] > prices[i-1] {
            profit += prices[i] - prices[i-1]
        }
    }
    return profit
}

func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }

    // 初始化：第0天持有股票，现金为 -prices[0]；不持有则为0[reference:12]
    hold := -prices[0]
    cash := 0

    for i := 1; i < len(prices); i++ {
        // 计算新的状态，注意顺序
        newCash := max(cash, hold+prices[i])
        newHold := max(hold, cash-prices[i])
        cash, hold = newCash, newHold
    }

    // 最终最大利润一定是不持有股票的状态
    return cash
}
