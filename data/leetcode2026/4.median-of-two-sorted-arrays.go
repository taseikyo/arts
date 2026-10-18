/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-18 18:13:34
 * @link    github.com/taseikyo
 */

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    total := len(nums1) + len(nums2)
    if total%2 == 1 {
        // 奇数长度，返回第 (total/2 + 1) 小的数
        return float64(findKth(nums1, nums2, total/2+1))
    }
    // 偶数长度，返回第 (total/2) 和 (total/2 + 1) 小数的平均值
    left := findKth(nums1, nums2, total/2)
    right := findKth(nums1, nums2, total/2+1)
    return float64(left+right) / 2.0
}

// findKth 寻找两个有序数组中第 k 小的元素 (k 从 1 开始)
func findKth(nums1, nums2 []int, k int) int {
    idx1, idx2 := 0, 0
    for {
        // 边界情况
        if idx1 == len(nums1) {
            return nums2[idx2+k-1]
        }
        if idx2 == len(nums2) {
            return nums1[idx1+k-1]
        }
        if k == 1 {
            return min(nums1[idx1], nums2[idx2])
        }

        // 正常二分查找
        half := k / 2
        // 防止越界，取实际能取的步数
        newIdx1 := min(idx1+half, len(nums1)) - 1
        newIdx2 := min(idx2+half, len(nums2)) - 1
        pivot1, pivot2 := nums1[newIdx1], nums2[newIdx2]

        if pivot1 <= pivot2 {
            // 排除 nums1 中 idx1 到 newIdx1 的部分
            k -= (newIdx1 - idx1 + 1)
            idx1 = newIdx1 + 1
        } else {
            // 排除 nums2 中 idx2 到 newIdx2 的部分
            k -= (newIdx2 - idx2 + 1)
            idx2 = newIdx2 + 1
        }
    }
}
