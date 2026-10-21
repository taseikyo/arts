/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-19 23:47:13
 * @link    github.com/taseikyo
 */

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

	// 如果辅助栈为空，或 val 小于等于辅助栈栈顶，则压入 val
	if len(this.minStack) == 0 || val <= this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, val)
	} else {
		// 否则，将当前最小值（辅助栈栈顶）再次压入，保持同步
		this.minStack = append(this.minStack, this.minStack[len(this.minStack)-1])
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	// 主栈和辅助栈同时弹出
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.minStack) == 0 {
		return 0
	}
	return this.minStack[len(this.minStack)-1]
}
