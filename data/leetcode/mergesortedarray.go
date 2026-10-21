/*
* @Date:   2025-03-04 00:34:45
* @Author: Lewis Tian (taseikyo@gmail.com)
*/
func merge(nums1 []int, m int, nums2 []int, n int)  {
    if n == 0 {
        return
    }

    num := make([]int, m+n)
    i, j, k := 0, 0, 0
    for i < m || j < n {
        if i < m && j < n {
            if nums1[i] < nums2[j] {
                num[k] = nums1[i]
                i++
            } else {
                num[k] = nums2[j]
                j++
            }
        } else if i < m {
            num[k] = nums1[i]
            i++
        } else {
            num[k] = nums2[j]
            j++
        }
        k++
    }

    i = 0
    for i < m+n {
        nums1[i] = num[i]
        i++
    }
}
