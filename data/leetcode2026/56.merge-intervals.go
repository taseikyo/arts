/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-06 22:30:06
 * @link    github.com/taseikyo
 */


func merge(intervals [][]int) ([][]int) {
    slices.SortFunc(intervals, func(p, q []int) int { return p[0] - q[0] }) // 按照左端点从小到大排序

    res := make([][]int, 0)
    res = append(res, intervals[0])
    for _, interval := range intervals {
        length := len(res)
        if res[length-1][1] >= interval[0] {
            res[length-1][1] = max(res[length-1][1], interval[1])
        } else {
            res = append(res, interval)
        }
    }


    return res
}
