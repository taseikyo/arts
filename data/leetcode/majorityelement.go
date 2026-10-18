/*
* @Date:   2025-03-07 00:22:14
* @Author: Lewis Tian (taseikyo@gmail.com)
*/
func majorityElement(nums []int) int {
    res, count := 0, 0
    for _, num := range nums {
        if count == 0 {
            res = num
        }
        if res == num {
            count++
        } else {
            count--
        }
    }

    return res
}
