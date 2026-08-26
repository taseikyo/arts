/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-22 00:47:03
 * @link    github.com/taseikyo
 */

import "container/heap"

// 定义一个 IntHeap 类型，实现 heap.Interface 接口
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
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

func inventoryManagement(stock []int, cnt int) []int {
	h := &IntHeap{}
	heap.Init(h)

	for _, num := range stock {
		heap.Push(h, num)
		// 保持堆的大小为 k
		if h.Len() > cnt {
			heap.Pop(h)
		}
	}

	return (*h)
}
