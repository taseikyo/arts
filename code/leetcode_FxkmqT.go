/**
 * @date    2020-12-30 10:07:01
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @link 	github.com/taseikyo
 */

package main

import "fmt"

func solve(n int) (ans int) {
    cards := make([]int, 2*n)
    var dfs func(num int)
    dfs = func(num int) {
        if num > n {
            ans++
            return
        }
        for i := 0; i+num+1 < n*2; i++ {
            // 枚举两张数字为 num 的牌的放入空位，它们的间隔为 num
            if cards[i] == 0 && cards[i+num+1] == 0 {
                cards[i], cards[i+num+1] = num, num
                dfs(num + 1)
                cards[i], cards[i+num+1] = 0, 0
            }
        }
    }
    dfs(1)
    return
}