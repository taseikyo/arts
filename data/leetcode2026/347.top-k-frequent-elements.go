/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-20 00:49:34
 * @link    github.com/taseikyo
 */

import "container/heap"

// Pair 存储元素及其频率
type Pair struct {
	Num  int
	Freq int
}

// MinHeap 实现一个小顶堆，基于频率排序
type MinHeap []Pair

func (h MinHeap) Len() int { return len(h) }

// 小顶堆：频率低的排在前面
func (h MinHeap) Less(i, j int) bool { return h[i].Freq < h[j].Freq }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
	// 1. 统计频率
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// 2. 使用小顶堆维护频率前 k 高的元素
	h := &MinHeap{}
	heap.Init(h)
	for num, freq := range freqMap {
		heap.Push(h, Pair{Num: num, Freq: freq})
		if h.Len() > k {
			heap.Pop(h) // 弹出频率最小的元素
		}
	}

	// 3. 提取结果
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = heap.Pop(h).(Pair).Num
	}
	return result
}
