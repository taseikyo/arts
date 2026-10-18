/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-07-13 13:20:39
 * @link    github.com/taseikyo
 */

func majorityElement(nums []int) int {
    sort.Ints(nums)
    return nums[len(nums)/2]
}

func majorityElement(nums []int) int {
    var res, count int
    for _, v := range nums {
        if count == 0 {
            res = v
        }
        if res == v {
            count++
        } else {
            count--
        }
    }

    return res
}
