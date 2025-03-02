/*
* @Date:   2025-03-02 18:08:08
* @Author: Lewis Tian (taseikyo@gmail.com)
*/

func twoSum(nums []int, target int) []int {
    cache := make(map[int]int)
    for idx, val := range nums {
        if preIdx, ok := cache[target-val]; ok {
            return []int{preIdx, idx}
        }
        cache[val] = idx
    }

    return []int{}
}
