/*
* @Date:   2025-03-07 00:11:46
* @Author: Lewis Tian (taseikyo@gmail.com)
*/
func singleNumber(nums []int) int {
    x := 0
    for _, val := range nums {
        x = x^val
    }

    return x
}
