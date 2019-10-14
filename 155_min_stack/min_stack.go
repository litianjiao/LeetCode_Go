package lc0155

type MinStack struct {
	stack []unit
}

type unit struct {
	val, min int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}

}

func (this *MinStack) Push(x int) {
	min := x
	if len(this.stack) > 0 && x > this.GetMin() {
		min = this.GetMin()
	}
	this.stack = append(this.stack, unit{val: x, min: min})
}

func (this *MinStack) Pop() {

	this.stack = this.stack[:len(this.stack)-1]
}

// 获取栈顶元素
func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1].val
}

func (this *MinStack) GetMin() int {

	return this.stack[len(this.stack)-1].min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
