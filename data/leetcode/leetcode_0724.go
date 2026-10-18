/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2021-01-28 12:21:14
 * @link    github.com/taseikyo
 */

/**
 * 题目不难，先拿到总和，然后遍历比较即可
 */

func pivotIndex(nums []int) int {
    total := 0
    for _, v := range nums {
        total += v
    }
    sum := 0
    for i, v := range nums {
        if 2*sum+v == total {
            return i
        }
        sum += v
    }
    return -1
}