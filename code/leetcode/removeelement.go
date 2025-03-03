/*
* @Date:   2025-03-04 00:51:43
* @Author: Lewis Tian (taseikyo@gmail.com)
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
