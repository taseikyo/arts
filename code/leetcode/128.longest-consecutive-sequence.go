/*
* @Date:   2025-04-10 23:10:13
* @Author: Lewis Tian (taseikyo@gmail.com)
*/
func longestConsecutive(nums []int) int {
    dict := make(map[int]int)
    for _, num := range nums {
        if _, ok := dict[num]; !ok {
            prev := dict[num-1]
            next := dict[num+1]

            dict[num] += prev+next+1
            dict[num-prev] = dict[num]
            dict[num+next] = dict[num]
        }
    }

    res := 0
    for _, v := range dict {
        if v > res {
            res = v
        }
    }
    return res
}

func longestConsecutive2(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    sort.Ints(nums)
    res := 1
    cur := 1
    for j := 1; j < len(nums); j++ {
        switch nums[j] - nums[j-1] {
        case 1:
            cur++
            res = max(res, cur)
        case 0:
        default:
            cur = 1
        }
    }
    return res
}