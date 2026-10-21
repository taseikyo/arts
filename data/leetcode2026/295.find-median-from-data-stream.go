/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-20 00:58:02
 * @link    github.com/taseikyo
 */

import (
	"container/heap"
	"sort"
)

// hp 定义了一个最小堆
// 它通过内嵌 sort.IntSlice 来获得 Len() 和 Less() 方法
// 默认的 Less 方法实现的是升序，即小根堆[reference:8]
type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

type MedianFinder struct {
	left  hp // 存储较小的一半，通过存储负数实现为最大堆[reference:9]
	right hp // 存储较大的一半，原生最小堆
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

// AddNum 添加一个数字到数据结构中
func (this *MedianFinder) AddNum(num int) {
	// 1. 当两堆等长时，新元素加入后应使 left 比 right 多 1 个
	//    先将元素加入 right，再从 right 弹出一个最小元素加入 left[reference:10]
	if this.left.Len() == this.right.Len() {
		heap.Push(&this.right, num)
		minFromRight := heap.Pop(&this.right).(int)
		heap.Push(&this.left, -minFromRight) // 存入负数，实现大根堆[reference:11]
	} else {
		// 2. 当 left 比 right 多 1 个时，新元素加入后应使两堆等长
		//    先将元素（取反后）加入 left，再从 left 弹出一个最大元素加入 right[reference:12]
		heap.Push(&this.left, -num)                // 存入负数
		maxFromLeft := -heap.Pop(&this.left).(int) // 弹出时取反，得到真实值
		heap.Push(&this.right, maxFromLeft)
	}
}

// FindMedian 返回当前所有元素的中位数
func (this *MedianFinder) FindMedian() float64 {
	if this.left.Len() == this.right.Len() {
		// 偶数个元素：中位数为两个堆顶的平均值[reference:13]
		leftTop := -this.left.IntSlice[0] // left 存储的是负数，需要取反
		rightTop := this.right.IntSlice[0]
		return float64(leftTop+rightTop) / 2.0
	}
	// 奇数个元素：中位数为 left 的堆顶[reference:14]
	// left 存储的是负数，需要取反
	return float64(-this.left.IntSlice[0])
}
