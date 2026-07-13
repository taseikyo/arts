/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-07-13 12:56:19
 * @link    github.com/taseikyo
 */

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 只有 nums1
	if n == 0 {
		return
	}
	// 只有 nums2
	if m == 0 {
		for idx := range nums2 {
			nums1[idx] = nums2[idx]
		}
		return
	}
	i, j, k := m-1, n-1, m+n-1
	for i >= 0 || j >= 0 {
		if i >= 0 && j >= 0 {
			if nums1[i] > nums2[j] {
				nums1[k] = nums1[i]
				i--
			} else {
				nums1[k] = nums2[j]
				j--
			}
		} else if j >= 0 {
			nums1[k] = nums2[j]
			j--
		} else {
			break
		}
		k--
	}
}