/*
* @Date:   2025-07-10 00:18:13
* @Author: Lewis Tian (taseikyo@gmail.com)
*/
func searchInsert(nums []int, target int) int{
    left, right := 0, len(nums) -1
    for left <= right {
        mid := (left + right) / 2
        if nums[mid] > target {
            right = mid -1
        } else if nums[mid] < target {
            left = mid +1
        } else {
            return mid
        }
    }
    
    return left
}
