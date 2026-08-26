/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-23 00:31:08
 * @link    github.com/taseikyo
 */

func reversePairs(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	// 辅助数组，用于归并时临时存储排序结果
	tmp := make([]int, n)

	// 定义递归函数：对 nums[l..r] 进行归并排序，并返回该区间内的逆序对数
	var mergeSort func(l, r int) int
	mergeSort = func(l, r int) int {
		// 1. 递归终止：区间内只有一个元素，没有逆序对
		if l >= r {
			return 0
		}

		// 2. 分：将区间一分为二，递归计算左右两部分的逆序对
		mid := (l + r) >> 1
		ans := mergeSort(l, mid) + mergeSort(mid+1, r)

		// 3. 治：合并两个有序子数组，并统计跨左右的逆序对
		i, j, k := l, mid+1, 0
		for i <= mid && j <= r {
			if nums[i] <= nums[j] {
				// 左元素 <= 右元素，不构成逆序对
				tmp[k] = nums[i]
				i++
			} else {
				// 左元素 > 右元素，构成逆序对
				// 左半部分从 i 到 mid 的所有元素都 > nums[j]
				ans += mid - i + 1
				tmp[k] = nums[j]
				j++
			}
			k++
		}

		// 将剩余元素复制到 tmp
		for ; i <= mid; i, k = i+1, k+1 {
			tmp[k] = nums[i]
		}
		for ; j <= r; j, k = j+1, k+1 {
			tmp[k] = nums[j]
		}

		// 将 tmp 中排好序的结果写回 nums[l..r]
		for i = l; i <= r; i++ {
			nums[i] = tmp[i-l]
		}

		return ans
	}

	return mergeSort(0, n-1)
}
