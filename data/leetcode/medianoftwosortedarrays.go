/*
* @Date:   2025-03-02 19:14:05
* @Author: Lewis Tian (taseikyo@gmail.com)
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
    if len1 == 0 {
		return mid(nums2)
	}
    if len2 == 0 {
        return mid(nums1)
    }
    nums := make([]int, len1+len2)
    i, j, k := 0, 0, 0
    for i < len1 || j < len2 {
        if i < len1 && j < len2 {
            if nums1[i] < nums2[j] {
                nums[k] = nums1[i]
                i++
            } else {
                nums[k] = nums2[j]
                j++
            }
        } else if i < len1 {
            nums[k] = nums1[i]
            i++
        } else {
            nums[k] = nums2[j]
            j++
        }
        k++
    }

    return mid(nums)
}

func mid(nums []int) float64 {
	numLen := len(nums)
	if numLen == 0 {
		return 0
	}

	if numLen%2 == 0 {
		return float64(nums[numLen/2-1] + nums[numLen/2]) / 2
	}

	return float64(nums[numLen/2])
}
