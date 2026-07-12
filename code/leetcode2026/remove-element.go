/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-07-12 17:01:05
 * @link    github.com/taseikyo
 */

func removeElement(nums []int, val int) int {
    numLen := len(nums)
    if numLen == 0 {
        return 0
    }
    i, j := 0, numLen - 1
    count := 0
    for i <= j {
        if nums[i] == val {
            nums[i] = nums[j]
            j--
        } else {
            i++
            count++
        }
    }

    return count
}