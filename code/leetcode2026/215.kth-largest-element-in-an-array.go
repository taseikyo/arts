/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-20 00:38:52
 * @link    github.com/taseikyo
 */

import "math/rand"

func findKthLargest(nums []int, k int) int {
    // 为了确保平均时间复杂度为 O(n)，随机打乱数组
    // 可以避免在极端有序情况下退化为 O(n^2)
    rand.Shuffle(len(nums), func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })
    return quickSelect(nums, 0, len(nums)-1, k)
}

func quickSelect(nums []int, left, right, k int) int {
    // 1. 调用 partition，得到基准值的最终位置 index
    index := partition(nums, left, right)

    // 2. 判断基准值在降序数组中的排名
    // 因为 partition 将大的元素放在了左边，所以 index-left+1 就是基准值是第几大的
    rank := index - left + 1

    if rank == k {
        return nums[index]
    } else if rank > k {
        // 目标在左半部分
        return quickSelect(nums, left, index-1, k)
    } else {
        // 目标在右半部分，注意 k 要减去左半部分的大小
        return quickSelect(nums, index+1, right, k-rank)
    }
}

// partition 函数：将区间 [left, right] 划分为两部分
// 左边都大于等于基准值，右边都小于基准值
// 返回基准值最终的索引位置
func partition(nums []int, left, right int) int {
    // 随机选择一个元素作为基准值，并将其交换到最右边
    pivotIdx := left + rand.Intn(right-left+1)
    nums[pivotIdx], nums[right] = nums[right], nums[pivotIdx]

    pivot := nums[right]
    i := left // i 指向下一个大于等于 pivot 的元素应该放置的位置

    for j := left; j < right; j++ {
        // 将大于等于 pivot 的元素交换到左边
        if nums[j] >= pivot {
            nums[i], nums[j] = nums[j], nums[i]
            i++
        }
    }
    // 最后将 pivot 交换到正确的位置
    nums[i], nums[right] = nums[right], nums[i]
    return i
}

import "container/heap"

// 定义一个 IntHeap 类型，实现 heap.Interface 接口
// 这是一个小顶堆 (Min-Heap)
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // 小顶堆
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push 和 Pop 方法需要使用指针接收者，因为它们会修改 slice 的长度
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &IntHeap{}
	heap.Init(h)

	for _, num := range nums {
		heap.Push(h, num)
		// 保持堆的大小为 k
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// 堆顶元素就是第 k 大的元素
	return (*h)[0]
}
