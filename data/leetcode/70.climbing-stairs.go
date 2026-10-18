/*
* @Date:   2025-04-08 22:34:36
* @Author: Lewis Tian (taseikyo@gmail.com)
*/

var cache = map[int]int{
    1: 1,
    2: 2,
}

func climbStairs(n int) int {
    if v, ok := cache[n]; ok {
        return v
    }
    cache[n] = climbStairs(n-1) + climbStairs(n-2)

    return cache[n]
}
