/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-24 15:46:05
 * @link    github.com/taseikyo
 */

type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (m *MinStack) Push(x int) {
	m.stack = append(m.stack, x)
	if len(m.minStack) == 0 || x < m.minStack[len(m.minStack)-1] {
		m.minStack = append(m.minStack, x)
	} else {
		m.minStack = append(m.minStack, m.minStack[len(m.minStack)-1])
	}
}

func (m *MinStack) Pop() {
	if len(m.stack) == 0 {
		return
	}
	m.stack = m.stack[:len(m.stack)-1]
	m.minStack = m.minStack[:len(m.minStack)-1]
}

func (m *MinStack) Top() int {
	if len(m.stack) == 0 {
		return 0
	}
	return m.stack[len(m.stack)-1]
}

func (m *MinStack) GetMin() int {
	if len(m.minStack) == 0 {
		return 0
	}
	return m.minStack[len(m.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
