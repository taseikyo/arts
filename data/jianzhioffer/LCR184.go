/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-24 16:11:33
 * @link    github.com/taseikyo
 */

type Checkout struct {
	queue []int // 1. 普通队列，存储所有元素
	deque []int // 2. 双端队列，用切片模拟，存储单调递减的“候选最大值”
}

func Constructor() Checkout {
	return Checkout{
		queue: []int{},
		deque: []int{},
	}
}

func (c *Checkout) Get_max() int {
	if len(c.deque) == 0 {
		return -1
	}
	return c.deque[0]
}

func (c *Checkout) Add(value int) {
	c.queue = append(c.queue, value)

	for len(c.deque) > 0 && c.deque[len(c.deque)-1] < value {
		c.deque = c.deque[:len(c.deque)-1]
	}
	c.deque = append(c.deque, value)
}

func (c *Checkout) Remove() int {
	if len(c.queue) == 0 {
		return -1
	}

	v := c.queue[0]
	c.queue = c.queue[1:]
	if len(c.deque) > 0 && c.deque[0] == v {
		c.deque = c.deque[1:]
	}

	return v
}

/**
 * Your Checkout object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get_max();
 * obj.Add(value);
 * param_3 := obj.Remove();
 */
